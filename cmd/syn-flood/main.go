package main

import (
	"github.com/bilalcaliskan/syn-flood/pkg/options"
	"github.com/dimiro1/banner"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func init() {
	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}

func main() {
	sfo := options.GetSynFloodOptions()
	log.Println(sfo)
	// raw.StartFlooding(sfo.DstIpStr, sfo.DstPort, sfo.PayloadLength)
}
