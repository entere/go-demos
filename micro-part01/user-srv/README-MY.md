
## 1、创建 srv 模板

先安装micro

```bash
$ go get github.com/micro/micro
$ cd $GOPATH/src/github.com/micro/micro
$ go build -o micro main.go
$ sudo cp micro /usr/local/bin/

```

然后用micro命令创建模板
```bash 
$ micro new micro-part01/user-srv --namespace=mu.micro.ci --alias=user --type=srv --gopath=false
```
## 2、安装启动 consul

```bash
$ brew install consul
$ consul agent -dev
$ consul agent -dev -ui # 带ui http://localhost:8500/ui


```

## 3、使用

```bash
$ make build
$ ./user-srv

```

启动成功，我们调用Service.QueryUserByName测试一下服务是否正常:

```bash
$ micro --registry=consul call mu.micro.ci.srv.user User.QueryUserByName '{"userName":"entere"}'
```