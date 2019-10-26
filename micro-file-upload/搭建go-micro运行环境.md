# 本章节的目的是让大家最快速搭建好go-micro环境

## 1、安装micro和go-micro

```bash
$ go get -u -v github.com/micro/micro
$ go get -u -v github.com/micro/go-micro
$ cd $GOPATH/src/github.com/micro/micro
$ go build -o micro main.go
$ sudo cp micro /usr/local/bin/
```

## 2、安装protobuf

```bash
# 源码安装

$ wget https://github.com/protocolbuffers/protobuf/releases/download/v3.6.1/protobuf-all-3.6.1.tar.gz
$ tar zxvf protobuf-all-3.6.1.tar.gz
$ cd protobuf-3.6.1/
$ ./autogen.sh
$ ./configure
$ make
$ make install
$ protoc -h

# 或者
$ git clone https://github.com/google/protobuf
$ cd protobuf
$ ./autogen.sh
$ ./configure
$ make
$ make check
$ sudo make install

```



## 3、安装 Go 版本的 protobuf 编译器

```bash
$ go get -u github.com/golang/protobuf/protoc-gen-go

```



## 4、安装 consul

```bash
$ brew install consul
$ consul agent -dev
$ consul agent -dev -ui # 带ui http://localhost:8500/ui

```