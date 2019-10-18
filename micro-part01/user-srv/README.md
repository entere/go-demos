# User Service

This is the User service

Generated with

```
micro new micro-part01/user-srv --namespace=mu.micro.ci --alias=user --type=srv --gopath=false
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: mu.micro.ci.srv.user
- Type: srv
- Alias: user

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
./user-srv
```

Build a docker image
```
make docker
```


## 启动consul
consul agent -dev -ui

http://localhost:8500/ui


启动成功，我们调用Service.QueryUserByName测试一下服务是否正常:
micro --registry=consul call mu.micro.ci.srv.user User.QueryUserByName '{"userName":"entere"}'