package main

import (
	"fileapi/controller"
	"net/http"
)

func main() {
	http.HandleFunc("/file/upload", controller.UploadFile)
	http.HandleFunc("/file/download", controller.DownloadFile)
	http.ListenAndServe(":8080", nil)
}
