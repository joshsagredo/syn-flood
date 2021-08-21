## Syn Flood

[![CI](https://github.com/bilalcaliskan/syn-flood/workflows/CI/badge.svg?event=push)](https://github.com/bilalcaliskan/syn-flood/actions?query=workflow%3ACI)
[![Docker pulls](https://img.shields.io/docker/pulls/bilalcaliskan/syn-flood)](https://hub.docker.com/r/bilalcaliskan/syn-flood/)
[![Go Report Card](https://goreportcard.com/badge/github.com/bilalcaliskan/syn-flood)](https://goreportcard.com/report/github.com/bilalcaliskan/syn-flood)

This project is developed with the objective of learning low level network operations with Golang. It starts a syn flood attack
with raw sockets. Do not use it with devil needs.

### Configuration
syn-flood can be customized with several command line arguments:
```
--dstIpStr                  Provide public ip of the destination
--dstPort                   Provide reachable port of the destination
--payloadLength             Provide payload length in bytes for each SYN packet
```

### Download

#### Binary
Binary can be downloaded from [Releases](https://github.com/bilalcaliskan/syn-flood/releases) page.

#### Docker
Docker image can be downloaded with below command:
```shell
$ docker run bilalcaliskan/syn-flood:latest
```

### Development
This project requires below tools while developing:
- [pre-commit](https://pre-commit.com/)
- [golangci-lint](https://golangci-lint.run/usage/install/) - required by [pre-commit](https://pre-commit.com/)
