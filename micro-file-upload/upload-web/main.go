package main

import (
	"net/http"

	"github.com/micro/go-micro/util/log"

	"github.com/entere/go-demos/micro-file-upload/upload-web/handler"
	"github.com/micro/go-micro/web"
)

func main() {
	// create new web service
	service := web.NewService(
		web.Name("io.github.entere.web.upload"),
		web.Version("latest"),
	)

	// initialise service
	if err := service.Init(); err != nil {
		log.Fatal(err)
	}

	// register html handler
	service.Handle("/", http.FileServer(http.Dir("html")))

	// register call handler
	service.HandleFunc("/upload/call", handler.UploadCall)
	service.HandleFunc("/upload/fileUpload", handler.FileUpload)

	// run service
	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
