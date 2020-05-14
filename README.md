# Service

[![GitHub Workflow Status](https://img.shields.io/github/workflow/status/gomicro/service/Build/master)](https://github.com/gomicro/service/actions?query=workflow%3ABuild+branch%3Amaster)
[![Go Reportcard](https://goreportcard.com/badge/github.com/gomicro/service)](https://goreportcard.com/report/github.com/gomicro/service)
[![License](https://img.shields.io/github/license/gomicro/service.svg)](https://github.com/gomicro/service/blob/master/LICENSE.md)
[![Release](https://img.shields.io/github/release/gomicro/service.svg)](https://github.com/gomicro/service/releases/latest)

Service is a bootstrapping generator for creating new microservices.

## Installing

### From Source

```
go get github.com/gomicro/service
cd $GOPATH/src/github.com/gomicro/service
make build install
```

### From Relases

See [latest releases](https://github.com/gomicro/service/releases/latest) for precompiled binaries.

## Usage

1. create new directory for desired service
2. run `service` within that directory
3. vendor your dependecies with your tool of choice
4. run and test the new service with `make run test`

## Versioning

The project will be versioned in accordance with [Semver 2.0.0](https://semver.org). See the [releases](https://github.com/gomicro/service/releases) section for the latest version. Until version 1.0.0 the project is considered to be unstable.

It is always highly recommended to vendor the version you are using.

## License
See [LICENSE.md](./LICENSE.md) for more information.
