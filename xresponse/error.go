package xresponse

type Error struct {
	Err        error
	IsRespUser bool
}

func NewError(err error, isRespUser ...bool) *Error {
	e := &Error{Err: err}
	if len(isRespUser) > 0 {
		e.IsRespUser = isRespUser[0]
	}
	return e
}
