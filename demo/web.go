package main

import (
	"net/http"
	"fmt"
	"go/demo/proto"
)

func main() {
	http.HandleFunc("/list/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("请求成功")
		arr2 := [3] int{1,3,5}
		writer.Write(proto.ResponseSuccess(arr2))

	})
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
