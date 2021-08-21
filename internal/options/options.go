package options

import (
	"github.com/spf13/pflag"
)

var synFloodOptions = &SynFloodOptions{}

func init() {
	synFloodOptions.addFlags(pflag.CommandLine)
	pflag.Parse()
}

// GetSynFloodOptions returns the pointer of SynFloodOptions
func GetSynFloodOptions() *SynFloodOptions {
	return synFloodOptions
}

// SynFloodOptions contains frequent command line and application options.
type SynFloodOptions struct {
	// DstIpStr is the public ip address of the target
	DstIpStr string
	// DstPort is the reachable port of the target
	DstPort int
	// PayloadLength is the payload length in bytes for each SYN packet
	PayloadLength int
}

func (sfo *SynFloodOptions) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&sfo.DstIpStr, "dstIpStr", "213.238.175.187", "Provide public ip of the destination")
	fs.IntVar(&sfo.DstPort, "dstPort", 443, "Provide reachable port of the destination")
	fs.IntVar(&sfo.PayloadLength, "payloadLength", 1400, "Provide payload length in bytes for each SYN packet")
}
