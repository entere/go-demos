# 生成 service 模板
## 用micro new 命令生成模板
```bash
$ cd go-demos
$ micro new micro-upload/upload-srv --type=srv --namespace=go.micro.ci --gopath=false

```

## 修改 proto/upload/upload.proto

```proto

syntax = "proto3";

package go.micro.ci.srv.upload;

service Upload {
	rpc Call(Request) returns (Response) {}
	rpc Stream(StreamingRequest) returns (stream StreamingResponse) {}
	rpc PingPong(stream Ping) returns (stream Pong) {}
	//文件上传
	rpc FileUpload(FileRequest) returns (Response) {}
}

// 文件上传Request
message FileRequest{
	//文件名
	string fileName = 1;
	//文件大小
	int64 size = 2;
	//文件内容
	bytes data = 3;
}
message Message {
	string say = 1;
}

message Request {
	string name = 1;
}

message Response {
	string msg = 1;
}

message StreamingRequest {
	int64 count = 1;
}

message StreamingResponse {
	int64 count = 1;
}

message Ping {
	int64 stroke = 1;
}

message Pong {
	int64 stroke = 1;
}


```


## 编译 upload.proto

```bash
$ protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/upload/upload.proto
```

#  生成 web 模板

## 用micro new 命令生成模板
```bash
$ cd go-demos
$ micro new micro-upload/upload-web --type=web --namespace=go.micro.ci --gopath=false
```


