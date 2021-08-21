package raw

import (
	"fmt"
	"math/rand"
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

		//macAddrs = append(macAddrs, fmt.Sprintf("%02x:%02x:%02x:%02x:%02x:%02x", buf[0], buf[1], buf[2], buf[3],
		//	buf[4], buf[5], buf[6], buf[7]))
	}

	return macAddrs
}
