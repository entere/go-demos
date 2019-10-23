# golang 微服务框架 go-micro 实战之文件上传



## 1、用micro new 命令生成 service 模板

```bash
$ cd go-demos
$ micro new micro-file-upload/upload-srv --type=srv --namespace=io.github.entere --gopath=false
```

## 2、修改.proto文件，定义文件上传的数据结构 

```bash
$ vi micro-file-upload/upload-srv/proto/upload/upload.proto

```

新增加如下数据据结构

```proto

syntax = "proto3";

package io.github.entere.srv.upload;

service Upload {
	rpc Call(Request) returns (Response) {}
	rpc Stream(StreamingRequest) returns (stream StreamingResponse) {}
	rpc PingPong(stream Ping) returns (stream Pong) {}

	// 文件上传
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

## 3、编译 upload.proto

```bash
$ cd micro-file-upload/upload-srv
$ protoc --proto_path=${GOPATH}/src:. --micro_out=. --go_out=. proto/upload/upload.proto
# 或者 make build
```

## 4、编写srv的 FileUpload 上传的方法

先把service 项目中所有 相对路径改成 github 的绝对路径 例如： micro-file-upload/upload-srv/proto/upload  修改为 github.com/entere/go-demos/micro-file-upload/upload-srv/proto/upload


在handler/upload.go中添加 FileUpload方法

```golang
package handler

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/micro/go-micro/util/log"

	upload "github.com/entere/go-demos/micro-file-upload/upload-srv/proto/upload"
)

// Upload ...
type Upload struct{}

// FileUpload 文件上传
func (e *Upload) FileUpload(ctx context.Context, req *upload.FileRequest, rsp *upload.Response) error {
	log.Log("Received Upload.FileUpload request")
	// 开始文件上传
	file, err := os.OpenFile("./"+req.FileName, os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		fmt.Println("文件打开失败")
		rsp.Msg = "文件上传失败：打开失败"
		return nil
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	nn, err := writer.Write(req.Data)
	if err != nil || nn != int(req.Size) {
		fmt.Println("文件写入失败", err, nn)
		rsp.Msg = "文件上传失败：写入失败"
		return nil
	}
	writer.Flush()
	rsp.Msg = "文件上传成功~~~~~"
	return nil
}

// Call is a single request handler called via client.Call or the generated client code
func (e *Upload) Call(ctx context.Context, req *upload.Request, rsp *upload.Response) error {
	log.Log("Received Upload.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

// Stream is a server side stream handler called via client.Stream or the generated client code
func (e *Upload) Stream(ctx context.Context, req *upload.StreamingRequest, stream upload.Upload_StreamStream) error {
	log.Logf("Received Upload.Stream request with count: %d", req.Count)

	for i := 0; i < int(req.Count); i++ {
		log.Logf("Responding: %d", i)
		if err := stream.Send(&upload.StreamingResponse{
			Count: int64(i),
		}); err != nil {
			return err
		}
	}

	return nil
}

// PingPong is a bidirectional stream handler called via client.Stream or the generated client code
func (e *Upload) PingPong(ctx context.Context, stream upload.Upload_PingPongStream) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}
		log.Logf("Got ping %v", req.Stroke)
		if err := stream.Send(&upload.Pong{Stroke: req.Stroke}); err != nil {
			return err
		}
	}
}

```

到此，上传的 service编写完成


## 5、用micro new 命令生成 web 模板

```bash
$ cd go-demos
$ micro new micro-file-upload/upload-web --type=web --namespace=io.github.entere --gopath=false
```

先把web 项目中所有 相对路径改成 github 的绝对路径 例如： micro-file-upload/upload-srv/proto/upload  修改为 github.com/entere/go-demos/micro-file-upload/upload-srv/proto/upload


## 6、添加web的 FileUpload方法，调用service

```bash
$ cd micro-file-upload/upload-web

$ vi handler/handler.go

```

修改handler.go如下

```golang
package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	upload "github.com/entere/go-demos/micro-file-upload/upload-srv/proto/upload"
	"github.com/micro/go-micro/client"
)

// UploadCall ...
func UploadCall(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	var request map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// call the backend service
	uploadClient := upload.NewUploadService("io.github.entere.srv.upload", client.DefaultClient)
	rsp, err := uploadClient.Call(context.TODO(), &upload.Request{
		Name: request["name"].(string),
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response := map[string]interface{}{
		"msg": rsp.Msg,
		"ref": time.Now().UnixNano(),
	}

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

// FileUpload ...web 调用 srv 文件上传
func FileUpload(w http.ResponseWriter, r *http.Request) {
	// decode the incoming request as json
	// var request map[string]interface{}
	// if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
	// 	http.Error(w, err.Error(), 500)
	// 	return
	// }
	response := map[string]interface{}{
		"ref": time.Now().UnixNano(),
	}
	file, header, err := r.FormFile("upload")
	if err != nil {
		fmt.Println("获取文件失败")
		response["msg"] = "获取文件失败"
		json.NewEncoder(w).Encode(response)
		return
	}
	defer file.Close()

	data := make([]byte, header.Size)
	// 读取上传文件的内容
	file.Read(data)

	// call the backend service
	uploadClient := upload.NewUploadService("io.github.entere.srv.upload", client.DefaultClient)
	rsp, err := uploadClient.FileUpload(context.TODO(), &upload.FileRequest{
		FileName: header.Filename,
		Size:     int64(header.Size),
		Data:     data,
	})
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}

	// we want to augment the response
	response["msg"] = rsp.Msg

	// encode and write the response as json
	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

```


## 7、修改 main.go 注册 fileUpload handler


```golang
service.HandleFunc("/upload/fileUpload", handler.FileUpload)

```

## 8、修改 html/index/html

```html
<!DOCTYPE html>
<html>
    <head>
        <title>Upload Web</title>
        <meta name="viewport" content="width=device-width, initial-scale=1">
        
    </head>
    <body>
        

        <form enctype="multipart/form-data" method="post" action="/upload/fileUpload">

            <input type="file" name="upload"/>
            <br>
            <button type="submit" name="btn">上传</button>
            
        </form>


       
    </body>
</html>

```


9、运行

```bash
$ cd go-demos/micro-file-upload/upload-srv
$ make build
$ ./upload-srv

$ cd go-demos/micro-file-upload/upload-web
$ make build
$ ./upload-web

```

在浏览器中运行 http://localhost:port 


10、结束