package mobile

import "fmt"

type Error struct {
	Code    int
	Message string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) String() string {
	return fmt.Sprintf("code=%d, message=%s", e.Code, e.Message)
}

func NewError(code int, message string) *Error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
