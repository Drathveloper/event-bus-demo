package error

import (
	"fmt"
)

type DomainErrorCode string

const (
	ItemNotFound       DomainErrorCode = "ITEM_NOT_FOUND"
	DatabaseError      DomainErrorCode = "DATABASE_ERROR"
	ConfigurationError DomainErrorCode = "CONFIGURATION_ERROR"
	CryptographicError DomainErrorCode = "CRYPTO_ERROR"
	GenericError       DomainErrorCode = "GENERIC_ERROR"
)

type DomainError interface {
	Error() string
	GetCode() DomainErrorCode
	GetMessage() string
}

type domainError struct {
	Code    DomainErrorCode
	Message string
}

func (e domainError) Error() string {
	return fmt.Sprintf("err: %s - %s", e.Code, e.Message)
}

func (e domainError) GetCode() DomainErrorCode {
	return e.Code
}

func (e domainError) GetMessage() string {
	return e.Message
}

func NewItemNotFoundError(message string) DomainError {
	return &domainError{
		Code:    ItemNotFound,
		Message: message,
	}
}

func NewDatabaseError(message string) DomainError {
	return &domainError{
		Code:    DatabaseError,
		Message: message,
	}
}

func NewConfigurationError(message string) DomainError {
	return &domainError{
		Code:    ConfigurationError,
		Message: message,
	}
}

func NewGenericError(message string) DomainError {
	return &domainError{
		Code:    GenericError,
		Message: message,
	}
}

func NewCryptoError(message string) DomainError {
	return &domainError{
		Code:    CryptographicError,
		Message: message,
	}
}
