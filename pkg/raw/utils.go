package raw

import "math/rand"

func getRandomPayload(length int) []byte {
	payload := make([]byte, length)
	rand.Read(payload)
	return payload
}

func getIps() []string {
	ips := make([]string, 0)
	ips = append(ips, "117.58.245.110")
	return ips
}

func getPorts() []uint16 {
	ports := make([]uint16, 0)
	for i := uint16(1024); i <= uint16(65535); i++ {
		ports = append(ports, i)
	}

	return ports
}