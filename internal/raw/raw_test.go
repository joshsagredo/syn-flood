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
			443, 1000, srcIps[rand.Intn(len(srcIps))], "93.184.216.34",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))]},
		{
			"100byte_ack", "ack", 10, srcPorts[rand.Intn(len(srcPorts))],
			443, 1000, srcIps[rand.Intn(len(srcIps))], "93.184.216.34",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))],
		},
		{
			"100byte_synack", "synAck", 10, srcPorts[rand.Intn(len(srcPorts))],
			443, 1000, srcIps[rand.Intn(len(srcIps))], "93.184.216.34",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))],
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			stopChan := make(chan struct{})
			stoppedChan := make(chan struct{})
			t.Logf("starting flood, caseName=%s, floodType=%s, floodMilliSeconds=%d\n", tc.name, tc.floodType, tc.floodMilliSeconds)
			go func(stopChan chan struct{}, dstIp string, dstPort int, payloadLength int, floodType string) {
				err := StartFlooding(stopChan, dstIp, dstPort, payloadLength, floodType)
				if err != nil {
					stoppedChan <- struct{}{}
					t.Errorf("an error occured on flooding process: %s\n", err.Error())
					return
				}
				stoppedChan <- struct{}{}
			}(stopChan, tc.dstIp, tc.dstPort, tc.payloadLength, tc.floodType)

			time.Sleep(50 * time.Millisecond)
			stopChan <- struct{}{}
			t.Logf("\nshouldStop channel received a signal, stopping\n")
			<-stoppedChan
		})
	}
}
