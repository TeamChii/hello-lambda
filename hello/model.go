package hello

import "github.com/TeamChii/hello-lambda/common"

type HelloRequest struct {
	Name string `json:"name"`
}

type HelloResponse struct {
	HeaderResp common.HeaderResp `json:"headerResp"`
	Msg        string            `json:"msg"`
}
