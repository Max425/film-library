package common

type ErrorType struct {
	t string
}

var (
	ErrNotFound   = ErrorType{"not found"}
	ErrInternal   = ErrorType{"internal error"}
	ErrBadRequest = ErrorType{"bad request"}
)

func (er *ErrorType) String() string {
	return er.t
}
