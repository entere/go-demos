# Upload Service

This is the Upload service

Generated with

```
micro new micro-file-upload/upload-web --namespace=io.github.entere --type=web
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: io.github.entere.web.upload
- Type: web
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
./upload-web
```

Build a docker image
```
make docker
```