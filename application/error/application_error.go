package error

import (
	"fmt"
	"net/http"
)

type ApplicationError interface {
	Error() string
	GetCode() int
	GetMessage() string
}

type applicationError struct {
	Code    int
	Message string
}

func (e applicationError) Error() string {
	return fmt.Sprintf("err: %d - %s", e.Code, e.Message)
}

func (e applicationError) GetCode() int {
	return e.Code
}

func (e applicationError) GetMessage() string {
	return e.Message
}

func NewBadRequestError(message string) ApplicationError {
	return &applicationError{
		Code:    http.StatusBadRequest,
		Message: message,
	}
}

func NewUnauthorizedError(message string) ApplicationError {
	return &applicationError{
		Code:    http.StatusUnauthorized,
		Message: message,
	}
}

func NewForbiddenError(message string) ApplicationError {
	return &applicationError{
		Code:    http.StatusForbidden,
		Message: message,
	}
}

func NewNotFoundError(message string) ApplicationError {
	return &applicationError{
		Code:    http.StatusNotFound,
		Message: message,
	}
}

func NewPreconditionFailedError(message string) ApplicationError {
	return &applicationError{
		Code:    http.StatusPreconditionFailed,
		Message: message,
	}
}

func NewInternalServerError(message string) ApplicationError {
	return &applicationError{
		Code:    http.StatusInternalServerError,
		Message: message,
	}
}
