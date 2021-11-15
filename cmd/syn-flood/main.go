package main

import (
	"github.com/bilalcaliskan/syn-flood/internal/options"
	"github.com/bilalcaliskan/syn-flood/internal/raw"
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
	if err := raw.StartFlooding(sfo.DstIpStr, sfo.DstPort, sfo.PayloadLength); err != nil {
		log.Fatalf("an error occured on flooding process: %s", err.Error())
	}
}
