# Upload Service

This is the Upload service

Generated with

```
micro new micro-file-upload/upload-srv --namespace=io.github.entere --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: io.github.entere.srv.upload
- Type: srv
- Alias: upload

## Dependencies

Micro services depend on service discovery. The default is multicast DNS, a zeroconf system.

In the event you need a resilient multi-host setup we recommend consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./upload-srv
```

Build a docker image
```
make docker
```