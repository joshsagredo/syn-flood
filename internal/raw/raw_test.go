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
		{"500byte_syn", "syn", 500, srcPorts[rand.Intn(len(srcPorts))], 443, 100,
			srcIps[rand.Intn(len(srcIps))], "213.238.175.187",
			macAddrs[rand.Intn(len(macAddrs))], macAddrs[rand.Intn(len(macAddrs))]},
		// TODO: other cases like floodType=ack, floodType=synAck. think a solution to "too many open files" error ;)
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ctx, cancel := context.WithTimeout(context.Background(), time.Duration(tc.floodMilliSeconds)*time.Millisecond)
			defer cancel()
			t.Logf("starting flood, caseName=%s, floodType=%s, floodMilliSeconds=%d\n", tc.name, tc.floodType, tc.floodMilliSeconds)
			go func() {
				err := StartFlooding(tc.dstIp, tc.dstPort, tc.payloadLength, tc.floodType)
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
				t.Logf("ending flood, caseName=%s, floodType=%s, floodMilliSeconds=%d\n", tc.name, tc.floodType, tc.floodMilliSeconds)
			}
		})
	}
}
