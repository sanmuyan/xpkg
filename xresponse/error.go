package xresponse

type Error struct {
	Err        error
	IsRespUser bool
	Code       HTTPCode
}

func NewError(err error, isRespUser ...bool) *Error {
	e := &Error{Err: err}
	if len(isRespUser) > 0 {
		e.IsRespUser = isRespUser[0]
	}
	return e
}

func (e *Error) WithCode(code HTTPCode) *Error {
	e.Code = code
	return e
}
