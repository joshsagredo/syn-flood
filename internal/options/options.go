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
	// Host is the public ip address or DNS of the target
	Host string
	// DstPort is the reachable port of the Host
	Port int
	// PayloadLength is the payload length in bytes for each SYN packet
	PayloadLength int
	// FloodType is the type of the attack type, usable values are syn, ack, synack
	FloodType string
}

func (sfo *SynFloodOptions) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&sfo.Host, "host", "213.238.175.187", "Provide public ip or DNS of the target")
	fs.IntVar(&sfo.Port, "port", 443, "Provide reachable port of the target")
	fs.IntVar(&sfo.PayloadLength, "payloadLength", 1400, "Provide payload length in bytes for each SYN packet")
	fs.StringVar(&sfo.FloodType, "floodType", "syn", "Provide the attack type. Proper values are: syn, ack, synack")
}
