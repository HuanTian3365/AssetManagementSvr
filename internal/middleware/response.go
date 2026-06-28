package middleware

import (
	"asset_management_svr/internal/model/common"
	"net/http"

	"github.com/gogf/gf/v2/net/ghttp"
)

func ResponseHandler(r *ghttp.Request) {
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		return
	}

	err := r.GetError()
	res := r.GetHandlerResponse()

	if err != nil {

		//code := http.StatusInternalServerError // 降级默认
		//if gc := gerror.Code(err); gc.Code() > 0 && gc != gcode.CodeUnknown {
		//	code = gc.Code() // 透传 logic/controller 抛的码
		//}
		//common.Json(r, code, err.Error(), nil)

		common.Json(r, http.StatusInternalServerError, err.Error(), nil)
		return
	}
	common.Json(r, 0, "", res)
}
