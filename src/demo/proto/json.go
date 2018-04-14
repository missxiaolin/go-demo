package proto

import (
	"encoding/json"
	"fmt"
)

type RequestACK struct {
	Service   string
	Method    string
	Arguments []interface{}
}

type ResponseACK struct {
	Success      bool
	Data         interface{}
	ErrorCode    uint32
	ErrorMessage string
}

func RequestBytes(bys []byte) *RequestACK {
	data := new(RequestACK)
	var _ = json.Unmarshal(bys, data)
	return data
}

func ResponseSuccess(data string) []byte {
	responseAck := make(map[string]interface{})

	responseAck["success"] = true
	responseAck["data"] = data

	fmt.Println(responseAck)

	r, _ := json.Marshal(responseAck)

	return r
}
