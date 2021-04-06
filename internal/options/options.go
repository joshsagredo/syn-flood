package options

var synFloodOptions = &SynFloodOptions{}

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
	// FloodDurationSeconds is the duration of the attack with seconds. Defaults to -1 which means
	FloodDurationSeconds int64
	// BannerFilePath is the relative path to the banner file
	BannerFilePath string
	// VerboseLog is the verbosity of the logging library
	VerboseLog bool
}
