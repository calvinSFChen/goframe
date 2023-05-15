package cerrors

import "github.com/gogf/gf/v2/errors/gcode"

type Code struct {
	code    int
	message string
	detail  CodeDetail
}
type CodeDetail struct {
	Code     string
	HttpCode int
}

func (c Code) CodeDetail() CodeDetail {
	return c.detail
}

func (c Code) Code() int {
	return c.code
}

func (c Code) Message() string {
	return c.message
}

func (c Code) Detail() interface{} {
	return c.detail
}

func New(httpCode int, code string, message string) gcode.Code {
	return Code{
		code:    httpCode,
		message: message,
		detail: CodeDetail{
			Code:     code,
			HttpCode: httpCode,
		},
	}
}

var (
	CodeSuccessNil = New(2001, "", "")
	CodeNil        = New(201, "", "操作失败")
	CodeSuccess    = New(200, "", "操作成功")
	CodeNotFound   = New(404, "Not Found", "Resource does not exist")
	CodeInternal   = New(500, "Internal Error", "An error occurred internally")
	CodeExpire     = New(499, "Internal Error", "登录已过期")
)
