package common

type ErrorType struct {
	t string
}

var (
	ErrNotFound           = ErrorType{"not found"}
	ErrInternal           = ErrorType{"internal error"}
	InvalidMailOrPassword = ErrorType{"invalid mail or password"}
	ErrBadRequest         = ErrorType{"bad request"}
)

func (er *ErrorType) String() string {
	return er.t
}

func (er *ErrorType) Error() string {
	return er.t
}
