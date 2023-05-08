package api

import "strconv"

type IResponse interface {
	GetCode() string
	GetMessage() string
	IsOk() bool
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

type Response2 struct {
	Code    int64       `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func (r Response) GetCode() string {
	return r.Code
}

func (r Response2) GetCode() string {
	return strconv.FormatInt(r.Code, 10)
}

func (r Response) GetMessage() string {
	return r.Message
}

func (r Response2) GetMessage() string {
	return r.Message
}

func (r Response) IsOk() bool {
	return r.Code == "0"
}

func (r Response2) IsOk() bool {
	return r.Code == 0
}
