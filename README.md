# Syn Flood

[![CI](https://github.com/bilalcaliskan/syn-flood/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/syn-flood/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/bilalcaliskan/syn-flood)](https://hub.docker.com/r/bilalcaliskan/syn-flood/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/syn-flood)](https://goreportcard.com/report/github.com/bilalcaliskan/syn-flood)
[![codecov](https://codecov.io/gh/bilalcaliskan/syn-flood/branch/master/graph/badge.svg)](https://codecov.io/gh/bilalcaliskan/syn-flood)
[![Release](https://img.shields.io/github/release/bilalcaliskan/syn-flood.svg)](https://github.com/bilalcaliskan/syn-flood/releases/latest)
[![Go version](https://img.shields.io/github/go-mod/go-version/bilalcaliskan/syn-flood)](https://github.com/bilalcaliskan/syn-flood)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

This project is developed with the objective of learning low level network operations with Golang. It starts a syn flood attack
with raw sockets.

> **Please do not use that tool with devil needs.**

## Prerequisites
You need root access to run syn-flood

## Configuration
syn-flood can be customized with several command line arguments:
```
--dstIpStr                  Provide public ip of the destination
--dstPort                   Provide reachable port of the destination
--payloadLength             Provide payload length in bytes for each SYN packet
```

## Download
### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/syn-flood/releases) page.

After then, you can simply run binary by providing required command line arguments:
```shell
$ sudo ./syn-flood --dstIpStr 10.0.0.100 --dstPort 443
```

### Docker
Docker image can be downloaded with below command:
```shell
$ docker run bilalcaliskan/syn-flood:latest
```

## Development
This project requires below tools while developing:
- [Golang 1.16](https://golang.org/doc/go1.16)
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)

## License
Apache License 2.0
