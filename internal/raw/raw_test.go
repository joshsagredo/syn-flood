package raw

import (
	"math/rand"
	"testing"
	"time"
)

func TestStartFlooding(t *testing.T) {
	srcIps := getIps()
	srcPorts := getPorts()
	macAddrs := getMacAddrs()
	cases := []struct {
		name, floodType                 string
		payloadLength, srcPort, dstPort int
		floodMilliSeconds               int64
		srcIp, dstIp                    string
		srcMacAddr, dstMacAddr          []byte
	}{
		{"100byte_syn", "syn", 10, srcPorts[rand.Intn(len(srcPorts))],
			443, 100, srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))]},
		{
			"100byte_ack", "ack", 10, srcPorts[rand.Intn(len(srcPorts))],
			443, 100, srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))],
		},
		{
			"100byte_synack", "synAck", 10, srcPorts[rand.Intn(len(srcPorts))],
			443, 100, srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))],
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			shouldStop := make(chan bool)
			t.Logf("starting flood, caseName=%s, floodType=%s, floodMilliSeconds=%d\n", tc.name, tc.floodType, tc.floodMilliSeconds)
			go func() {
				err := StartFlooding(shouldStop, tc.dstIp, tc.dstPort, tc.payloadLength, tc.floodType)
				if err != nil {
					t.Errorf("an error occured on flooding process: %s\n", err.Error())
					return
				}
			}()

			<-time.After(time.Duration(tc.floodMilliSeconds) * time.Millisecond)
			shouldStop <- true
			close(shouldStop)

			for {
				select {
				case <-shouldStop:
					t.Logf("\nshouldStop channel received a signal, stopping\n")
					return
				default:
					continue
				}
			}
		})
	}
}
