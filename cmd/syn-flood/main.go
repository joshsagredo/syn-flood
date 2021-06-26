package main

import (
	"github.com/bilalcaliskan/syn-flood/pkg/options"
	"github.com/bilalcaliskan/syn-flood/pkg/raw"
)

func main() {
	sfo := options.GetSynFloodOptions()
	raw.StartFlooding(sfo.DstIpStr, sfo.DstPort, sfo.PayloadLength)
}
