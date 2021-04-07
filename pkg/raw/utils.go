package raw

import "math/rand"

func getRandomPort() int {
	return 0
}

func getRandomPayload(length int) []byte {
	payload := make([]byte, length)
	rand.Read(payload)
	return payload
}