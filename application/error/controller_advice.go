package error

import (
	errorDomain "event-bus-demo/domain/error"
)

type ControllerAdvice interface {
	TranslateError(err errorDomain.DomainError) ApplicationError
}

type controllerAdvice struct {
}

func NewControllerAdvice() ControllerAdvice {
	return &controllerAdvice{}
}

func (advice *controllerAdvice) TranslateError(err errorDomain.DomainError) ApplicationError {
	if err == nil {
		return nil
	}
	switch err.GetCode() {
	case errorDomain.ItemNotFound:
		return NewNotFoundError(err.GetMessage())
	case errorDomain.DatabaseError:
		return NewInternalServerError("error while processing information")
	case errorDomain.CryptographicError:
		return NewInternalServerError("error while performing encryption/decryption")
	case errorDomain.GenericError:
		return NewInternalServerError("unhandled error in domain model")
	default:
		return NewInternalServerError("")
	}
}
