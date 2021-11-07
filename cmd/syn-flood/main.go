package main

import (
	"github.com/bilalcaliskan/syn-flood/internal/logging"
	"github.com/bilalcaliskan/syn-flood/internal/options"
	"github.com/bilalcaliskan/syn-flood/internal/raw"
	"github.com/dimiro1/banner"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"strings"
)

var logger *zap.Logger

func init() {
	logger = logging.GetLogger()

	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}

func main() {
	sfo := options.GetSynFloodOptions()
	if err := raw.StartFlooding(sfo.DstIpStr, sfo.DstPort, sfo.PayloadLength); err != nil {
		logger.Fatal("an error occured on flooding process", zap.String("error", err.Error()))
	}
}
