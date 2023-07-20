package servmsg

import (
	"hotgo/internal/library/network/tcp"
	"hotgo/internal/model/input/servmsgin"
)

// 一些例子.

// ExampleHelloReq 一个tcp请求例子
type ExampleHelloReq struct {
	Name string `json:"name" description:"名字"`
}

type ExampleHelloRes struct {
	tcp.ServerRes
	Data *servmsgin.ExampleHelloModel `json:"data,omitempty" description:"数据集"`
}

// ExampleRPCHelloReq 一个rpc请求例子
type ExampleRPCHelloReq struct {
	Name string `json:"name" description:"名字"`
}

type ExampleRPCHelloRes struct {
	tcp.ServerRes
	Data *servmsgin.ExampleHelloModel `json:"data,omitempty" description:"数据集"`
}
