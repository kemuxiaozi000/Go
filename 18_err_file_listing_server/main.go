package main

import (
	"net/http"

	"./filelisting"
)

func main() {
	http.HandleFunc("/list/", filelisting.HandleFileList)

	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}

}
