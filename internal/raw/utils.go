package raw

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"net"
	"regexp"
)

func getRandomPayload(length int) []byte {
	payload := make([]byte, length)
	rand.Read(payload)
	return payload
}

func getIps() []string {
	ips := make([]string, 0)
	for i := 0; i < 20; i++ {
		ips = append(ips, fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256),
			rand.Intn(256), rand.Intn(256)))
	}

	return ips
}

func getPorts() []int {
	ports := make([]int, 0)
	for i := 1024; i <= 65535; i++ {
		ports = append(ports, i)
	}

	return ports
}

func getMacAddrs() [][]byte {
	macAddrs := make([][]byte, 0)
	for i := 0; i <= 50; i++ {
		buf := make([]byte, 6)
		_, err := rand.Read(buf)
		if err != nil {
			fmt.Println("error:", err)
			continue
		}
		macAddrs = append(macAddrs, buf)
	}

	return macAddrs
}

func isDNS(host string) bool {
	var (
		res bool
		err error
	)

	if res, err = regexp.MatchString(DnsRegex, host); err != nil {
		log.Fatalf("a fatal error occured while matching provided --host with DNS regex: %s", err.Error())
	}

	return res
}

func isIP(host string) bool {
	var (
		res bool
		err error
	)

	if res, err = regexp.MatchString(IpRegex, host); err != nil {
		log.Fatalf("a fatal error occured while matching provided --host with IP regex: %s", err.Error())
	}

	return res
}

func resolveHost(host string) string {
	if !isIP(host) && isDNS(host) {
		log.Printf("%s is a DNS record, making DNS lookup\n", host)
		ipRecords, err := net.DefaultResolver.LookupIP(context.Background(), "ip4", host)
		if err != nil {
			log.Fatalf("an error occured on dns lookup: %s", err.Error())
		}

		log.Printf("dns lookup succeeded, found %s for %s\n", ipRecords[0].String(), host)
		host = ipRecords[0].String()
	}

	return host
}
