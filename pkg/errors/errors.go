package errors

type Error struct {
	Code int
	Msg  string
	Err  error
}

func New(code int, msg string) *Error {
	return &Error{Code: code, Msg: msg}
}

// Error return error with info
func (e *Error) Error() string {
	return e.Msg
}

// WithMsg with message
func (e *Error) WithMsg(msg string) *Error {
	e.Msg = msg
	return e
}

// WithError with original error
func (e *Error) WithError(err error) *Error {
	e.Err = err
	return e
}
