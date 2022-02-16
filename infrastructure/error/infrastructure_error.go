package error

import (
	"fmt"
)

type InfrastructureErrorCode string

const (
	ItemNotFound   InfrastructureErrorCode = "ITEM_NOT_FOUND"
	SQLError       InfrastructureErrorCode = "SQL_ERROR"
	ParseFileError InfrastructureErrorCode = "PARSE_FILE_ERROR"
	HashingError   InfrastructureErrorCode = "HASHING_ERROR"
)

type InfrastructureError interface {
	Error() string
	GetCode() InfrastructureErrorCode
	GetMessage() string
}

type infrastructureError struct {
	Code    InfrastructureErrorCode
	Message string
}

func (e infrastructureError) Error() string {
	return fmt.Sprintf("err: %s - %s", e.Code, e.Message)
}

func (e infrastructureError) GetCode() InfrastructureErrorCode {
	return e.Code
}

func (e infrastructureError) GetMessage() string {
	return e.Message
}

func NewItemNotFoundError(message string) InfrastructureError {
	return &infrastructureError{
		Code:    ItemNotFound,
		Message: message,
	}
}

func NewSQLError(message string) InfrastructureError {
	return &infrastructureError{
		Code:    SQLError,
		Message: message,
	}
}

func NewParseFileError(message string) InfrastructureError {
	return &infrastructureError{
		Code:    ParseFileError,
		Message: message,
	}
}

func NewHashingError(message string) InfrastructureError {
	return &infrastructureError{
		Code:    HashingError,
		Message: message,
	}
}
