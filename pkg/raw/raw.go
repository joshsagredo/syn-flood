package raw

import (
	"github.com/google/gopacket/layers"
	"net"
)

func CreatePacket(srcIp, dstIp net.IP) layers.IPv4 {
	return layers.IPv4{
		SrcIP: srcIp,
		DstIP: dstIp,
		Version: 4,
		TTL: 64,
	}
}