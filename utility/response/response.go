package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

type ResponseRes struct {
	Code int         `json:"code" description:"状态码"`
	Msg  string      `json:"msg" description:"操作信息"`
	Data interface{} `json:"data" description:"返回信息"`
}

func JsonExit(r *ghttp.Request, code int, msg string, data interface{}) {
	r.Response.WriteJson(ResponseRes{
		Code: code,
		Msg:  msg,
		Data: data,
	})
	r.Exit()
}
