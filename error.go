package gox

import (
	"database/sql"
	"fmt"
	"net/http"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ErrorString string

func (es ErrorString) Error() string {
	return string(es)
}

const (
	ErrNotExist ErrorString = "does not exist"
)

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (e *Error) Error() string {
	return e.Message
}

func NewError(code int, msgFormat string, args ...interface{}) *Error {
	message := fmt.Sprintf(msgFormat, args...)
	if len(message) == 0 {
		message = http.StatusText(code % 1000)
	}
	return &Error{
		Code:    code,
		Message: message,
	}
}

func InternalError(format string, args ...interface{}) *Error {
	return NewError(http.StatusInternalServerError, format, args...)
}

func BadRequest(format string, args ...interface{}) *Error {
	return NewError(http.StatusBadRequest, format, args...)
}

func Unauthorized(format string, args ...interface{}) *Error {
	return NewError(http.StatusUnauthorized, format, args...)
}

func Forbidden(format string, args ...interface{}) *Error {
	return NewError(http.StatusForbidden, format, args...)
}

func NotFound(format string, args ...interface{}) *Error {
	return NewError(http.StatusNotFound, format, args...)
}

func Conflict(format string, args ...interface{}) *Error {
	return NewError(http.StatusConflict, format, args...)
}

func ToStatusError(err error) error {
	if err == nil {
		return nil
	}

	err = Cause(err)

	// if err is status error, return directly
	_, ok := status.FromError(err)
	if ok {
		return err
	}

	if err == ErrNotExist || err == sql.ErrNoRows {
		return status.Error(codes.Code(http.StatusNotFound), err.Error())
	}

	switch v := err.(type) {
	case *Error:
		return status.Error(codes.Code(v.Code), v.Error())
	default:
		return status.Error(codes.Code(http.StatusInternalServerError), err.Error())
	}
}

func FromStatusError(err error) error {
	if err == nil {
		return nil
	}

	s, ok := status.FromError(err)
	if !ok {
		return err
	}

	if int(s.Code()) == http.StatusNotFound &&
		(s.Message() == ErrNotExist.Error() || s.Message() == sql.ErrNoRows.Error()) {
		return ErrNotExist
	}

	return NewError(int(s.Code()), s.Message())
}

func Cause(err error) error {
	for {
		if e, ok := err.(interface{ Unwrap() error }); ok {
			err = e.Unwrap()
		} else {
			return err
		}
	}
}
