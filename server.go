package main

import (
	"fmt"
	"net/http"
	"./download"
	"./upload"
)

func main() {
	http.HandleFunc("/upload", upload.Uploadhandle)
	http.HandleFunc("/download", download.Downloader)
	err := http.ListenAndServe(":22048", nil)
	if err != nil {
		fmt.Println(err)
	}
}
