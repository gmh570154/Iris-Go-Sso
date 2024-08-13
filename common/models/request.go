package models

import (
	"github.com/google/uuid"
	"github.com/kataras/iris/v12"
)

type Request struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

type Response struct {
	ID      uint64 `json:"id"`
	Message string `json:"message"`
}

type ResponseBase struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func GenSuccessData(data interface{}) *ResponseBase {
	return &ResponseBase{200, "", data}
}

func GenSuccessMsg(msg string) *ResponseBase {
	return &ResponseBase{200, msg, ""}
}

func GenFailedMsg(errMsg string) *ResponseBase {
	return &ResponseBase{400, errMsg, ""}
}

var xRequestIDHeaderKey = "glb_request_id"

var GeneratorGlbRequestId = func(ctx iris.Context) string {
	id := ctx.Request().Header.Get(xRequestIDHeaderKey)
	if id != "" {
		return id
	}

	id = ctx.GetHeader(xRequestIDHeaderKey)
	if id == "" {
		uid, err := uuid.NewRandom()
		if err != nil {
			ctx.StopWithStatus(500)
			return ""
		}
		id = uid.String()
	}
	ctx.Request().Header.Set(xRequestIDHeaderKey, "glb_request_"+id)
	return id
}
