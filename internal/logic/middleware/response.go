package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"goframe/utility/cerrors"
	"goframe/utility/response"
	"net/http"
)

func (m *sMiddleware) HandlerResponse(r *ghttp.Request) {
	r.Middleware.Next()
	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 {
		return
	}

	var (
		err  = r.GetError()
		data = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		// 没有指定错误码
		if code == gcode.CodeNil {
			code = gcode.WithCode(cerrors.CodeNil, nil)
		}
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg := http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			code = gcode.New(code.Code(), msg, code.Detail())
		} else {
			code = gcode.WithCode(cerrors.CodeSuccess, nil)
		}
	}
	// 非成功状态，不返回数据
	if code.Code() != 200 {
		data = nil
	}
	response.JsonExit(r, code.Code(), code.Message(), data)
}
