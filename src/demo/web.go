package main

import (
	"net/http"
	"fmt"
	"demo/proto"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("请求成功")
		proto.ResponseSuccess("s")
	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
