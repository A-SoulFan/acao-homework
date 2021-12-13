package apperrors

import "fmt"

const (
	ValidateError = iota + 1
	AuthenticationError
	ServiceError
)

type Error struct {
	err     error
	errType int
	message string
}

func (e Error) Error() string {
	if e.err != nil {
		e.message = fmt.Sprintf("%s\n%+v", e.message, e.err)
	}
	return e.message
}

func (e Error) Message() string {
	return e.message
}

func (e Error) ErrorType() int {
	return e.errType
}

func (e *Error) Wrap(err error) *Error {
	e.err = err
	return e
}

func NewError(message string, t int) *Error {
	return &Error{message: message, errType: t}
}

func NewValidateError(message string) *Error {
	return &Error{message: message, errType: ValidateError}
}

func NewAuthenticationError(message string) *Error {
	return &Error{message: message, errType: AuthenticationError}
}

func NewServiceError(message string) *Error {
	return &Error{message: message, errType: ServiceError}
}
