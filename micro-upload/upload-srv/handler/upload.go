package handler

import (
	"bufio"
	"context"
	"fmt"
	"os"

	"github.com/micro/go-micro/util/log"

	upload "github.com/entere/go-demos/micro-upload/upload-srv/proto/upload"
)

// Upload ...
type Upload struct{}

// Call is a single request handler called via client.Call or the generated client code
func (e *Upload) Call(ctx context.Context, req *upload.Request, rsp *upload.Response) error {
	log.Log("Received Upload.Call request")
	rsp.Msg = "Hello " + req.Name
	return nil
}

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
