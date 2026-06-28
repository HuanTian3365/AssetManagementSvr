package common

import "github.com/gogf/gf/v2/net/ghttp"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func Json(r *ghttp.Request, code int, message string, data interface{}) {
	r.Response.WriteJson(Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

func JsonExit(r *ghttp.Request, code int, message string, data interface{}) {
	Json(r, code, message, data)
	r.ExitAll()
}
