package raw

import (
	"context"
	"fmt"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"math/rand"
	"net"
	"regexp"
)

// getRandomPayload returns a byte slice to spoof ip packets with random payload in specified length
func getRandomPayload(length int) []byte {
	payload := make([]byte, length)
	rand.Read(payload)
	return payload
}

// getIps returns a string slice to spoof ip packets with dummy source ip addresses
func getIps() []string {
	ips := make([]string, 0)
	for i := 0; i < 20; i++ {
		ips = append(ips, fmt.Sprintf("%d.%d.%d.%d", rand.Intn(256), rand.Intn(256),
			rand.Intn(256), rand.Intn(256)))
	}

	return ips
}

// getPorts returns an int slice to spoof ip packets with dummy source ports
func getPorts() []int {
	ports := make([]int, 0)
	for i := 1024; i <= 65535; i++ {
		ports = append(ports, i)
	}

	return ports
}

// getMacAddrs returns a byte slice to spoof ip packets with dummy MAC addresses
func getMacAddrs() [][]byte {
	macAddrs := make([][]byte, 0)
	for i := 0; i <= 50; i++ {
		buf := make([]byte, 6)
		_, err := rand.Read(buf)
		if err != nil {
			continue
		}
		macAddrs = append(macAddrs, buf)
	}

	return macAddrs
}

// isDNS returns a boolean which indicates host parameter is a DNS record or not
func isDNS(host string) bool {
	var (
		res bool
		err error
	)

	if res, err = regexp.MatchString(DnsRegex, host); err != nil {
		logger.Fatal("fatal error occurred while matching provided --host flag with DNS regex", zap.String("host", host),
			zap.String("regex", DnsRegex), zap.String("error", err.Error()))
	}

	return res
}

// isIP returns a boolean which indicates host parameter is an IP address or not
func isIP(host string) bool {
	var (
		res bool
		err error
	)

	if res, err = regexp.MatchString(IpRegex, host); err != nil {
		logger.Fatal("fatal error occurred while matching provided --host flag with IP regex", zap.String("host", host),
			zap.String("regex", IpRegex), zap.String("error", err.Error()))
	}

	return res
}

// resolveHost function gets a string and returns the ip address while deciding it is an ip address already or DNS record
func resolveHost(host string) (string, error) {
	if !isIP(host) && isDNS(host) {
		logger.Debug("already a DNS record provided, making DNS lookup", zap.String("host", host))
		ipRecords, err := net.DefaultResolver.LookupIP(context.Background(), "ip4", host)
		if err != nil {
			return "", errors.Wrapf(err, "an error occured on DNS lookup for %s", host)
		}

		logger.Debug("DNS lookup succeeded", zap.String("DNS", host), zap.String("IP", ipRecords[0].String()))
		host = ipRecords[0].String()
	} else {
		logger.Debug("already an IP address, skipping DNS resolution", zap.String("host", host))
	}

	return host, nil
}
