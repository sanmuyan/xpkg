package xresponse

// 封装 HTTP 框架的返回值

type Framework interface {
	SetFramework(*Response)
}

type HTTPCode int

const (
	HttpOk                  HTTPCode = 200
	HttpBadRequest          HTTPCode = 400
	HttpUnauthorized        HTTPCode = 401
	HttpForbidden           HTTPCode = 413
	HttpInternalServerError HTTPCode = 500
)

func (m HTTPCode) GetMessage() string {
	switch m {
	case HttpOk:
		return "操作成功"
	case HttpBadRequest:
		return "数据错误"
	case HttpUnauthorized:
		return "身份验证错误"
	case HttpForbidden:
		return "无权限访问"
	case HttpInternalServerError:
		return "服务器内部错误"
	}
	return ""
}

type Response struct {
	Success bool     `json:"success"`
	Code    HTTPCode `json:"code"`
	Message string   `json:"message,omitempty"`
	Data    any      `json:"data,omitempty"`
}

func (r *Response) defaultSet() {
	if r.Code < 600 {
		r.Code += 1000
	}
}

func (r *Response) Ok() *Response {
	code := HttpOk
	r.Code = code
	r.Success = true
	r.Message = code.GetMessage()
	return r
}

func (r *Response) Fail(code HTTPCode) *Response {
	r.Code = code
	r.Success = false
	r.Message = code.GetMessage()
	return r
}

func (r *Response) WithData(data any) *Response {
	r.Data = data
	return r
}

func (r *Response) WithMsg(msg string) *Response {
	r.Message = msg
	return r
}

func (r *Response) WithError(err *Error) *Response {
	if err == nil {
		return r
	}
	if err.IsRespUser {
		r.Message = err.Err.Error()
	}
	return r
}

func (r *Response) FailWithError(err *Error, code ...HTTPCode) *Response {
	if len(code) > 0 {
		r.Fail(code[0])
	} else {
		r.Fail(HttpInternalServerError)
	}
	if err.Code > 0 {
		r.Fail(err.Code)
		return r.WithError(err)
	}
	return r.WithError(err)
}

func (r *Response) Response(rf Framework) {
	r.defaultSet()
	rf.SetFramework(r)
}

func NewResponse() *Response {
	return &Response{
		Success: false,
		Code:    HttpOk,
		Message: "",
		Data:    nil,
	}
}
