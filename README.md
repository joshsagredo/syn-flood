# Syn Flood

[![CI](https://github.com/bilalcaliskan/syn-flood/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/syn-flood/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/bilalcaliskan/syn-flood)](https://hub.docker.com/r/bilalcaliskan/syn-flood/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/syn-flood)](https://goreportcard.com/report/github.com/bilalcaliskan/syn-flood)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_syn-flood&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_syn-flood)
[![Reliability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_syn-flood&metric=reliability_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_syn-flood)
[![Maintainability Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_syn-flood&metric=sqale_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_syn-flood)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_syn-flood&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_syn-flood)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=bilalcaliskan_syn-flood&metric=coverage)](https://sonarcloud.io/summary/new_code?id=bilalcaliskan_syn-flood)
[![Release](https://img.shields.io/github/release/bilalcaliskan/syn-flood.svg)](https://github.com/bilalcaliskan/syn-flood/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/bilalcaliskan/syn-flood)](https://github.com/bilalcaliskan/syn-flood)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This project is developed with the objective of learning low level network operations with Golang. It starts a syn flood attack
with raw sockets.

## Legal Disclaimer
This tool is created for the sole purpose of security awareness and education, it should not be used against systems
that you do not have permission to test/attack. The author is not responsible for misuse or for any damage that you
may cause. You agree that you use this software at your own risk.

## Prerequisites
You need root access to run syn-flood

syn-flood needs lots of open file descriptors while running so we need to increase it first. You can increase it like below
temporarily. That works for both Macos and Linux:

```shell
$ sudo ulimit -S -n 2048000
$ sudo syn-flood --host foo.example.com --port 443 --floodType syn
```

If you still get **"too many open files"** error, try increasing the value that passed to first command.

## Configuration
syn-flood can be customized with several command line arguments:
```
Flags:
      --floodDurationSeconds int   Provide the duration of the attack in seconds, -1 for no limit, defaults to -1 (default -1)
      --floodType string           Provide the attack type. Proper values are: syn, ack, synack (default "syn")
  -h, --help                       help for syn-flood
      --host string                Provide public ip or DNS of the target (default "213.238.175.187")
      --payloadLength int          Provide payload length in bytes for each SYN packet (default 1400)
      --port int                   Provide reachable port of the target (default 443)
  -v, --verbose                    verbose output of the logging library (default false)
      --version                    version for syn-flood
```

> To be able to run **syn-flood** with unlimited time range, you should also increase your operating system open file
> limits, you can refer [here](https://www.tecmint.com/increase-set-open-file-limits-in-linux/) about how to do that.

## Download
### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/syn-flood/releases) page.

After then, you can simply run binary by providing required command line arguments:
```shell
$ sudo ./syn-flood --host 10.0.0.100 --port 443
```

Or with DNS:
```shell
$ sudo ./syn-flood --host foo.example.com --port 443
```

### Homebrew
This project can be installed with [Homebrew](https://brew.sh/):
```shell
$ brew tap bilalcaliskan/tap
$ brew install bilalcaliskan/tap/syn-flood
$ sudo syn-flood --host foo.example.com --port 443 --floodType syn
```

### Docker
Docker image can be downloaded with below command:
```shell
$ docker run bilalcaliskan/syn-flood:latest
```

## Development
This project requires below tools while developing:
- [Golang 1.19](https://golang.org/doc/go1.19)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)
- [gocyclo](https://github.com/fzipp/gocyclo) - required by [pre-commit](https://pre-commit.com/)

After you installed [pre-commit](https://pre-commit.com/), simply run below command to prepare your development environment:
```shell
$ pre-commit install -c build/ci/.pre-commit-config.yaml
```

## References
- https://www.devdungeon.com/content/packet-capture-injection-and-analysis-gopacket
- https://www.programmersought.com/article/74831586115/
- https://github.com/rootVIII/gosynflood
- https://golangexample.com/repeatedly-send-crafted-tcp-syn-packets-with-raw-sockets/
- https://github.com/kdar/gorawtcpsyn/blob/master/main.go
- https://pkg.go.dev/github.com/google/gopacket
- https://github.com/david415/HoneyBadger/blob/021246788e58cedf88dee75ac5dbf7ae60e12514/packetSendTest.go
- mac spoofing -> https://github.com/google/gopacket/issues/153
- free proxies -> https://www.sslproxies.org/
- [What is an ACK flood DDoS attack? | Types of DDoS attacks](https://www.cloudflare.com/tr-tr/learning/ddos/what-is-an-ack-flood/)
- https://bariskoparmal.com/2021/08/22/spesifik-ddos-saldirilari-ve-saldiri-komutlari/
- https://github.blog/2016-07-12-syn-flood-mitigation-with-synsanity/

## License
Apache License 2.0
