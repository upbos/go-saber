package e

import (
	"github.com/pkg/errors"
)

type WrapError struct {
	Code    int
	Message string
	Data    interface{}
}

func (e *WrapError) Error() string {
	return e.Message
}

func New(code int) error {
	return &WrapError{Code: code}
}

func WithMessage(code int, message string) error {
	return &WrapError{Code: code, Message: message}
}

func WithData(code int, message string, data interface{}) error {
	return &WrapError{Code: code, Message: message, Data: data}
}

func Wrap(message string, err error) error {
	return errors.Wrap(err, message)
}

func UnWrap(err error) (*WrapError, bool) {
	v, ok := err.(*WrapError)
	return v, ok
}
