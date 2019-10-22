package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	upload "github.com/entere/go-demos/micro-upload/upload-srv/proto/upload"
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
	uploadClient := upload.NewUploadService("go.micro.ci.srv.upload", client.DefaultClient)
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
	uploadClient := upload.NewUploadService("go.micro.ci.srv.upload", client.DefaultClient)
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
