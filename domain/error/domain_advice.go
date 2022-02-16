package error

import (
	errorInfrastructure "event-bus-demo/infrastructure/error"
)

type DomainAdvice interface {
	TranslateError(err errorInfrastructure.InfrastructureError) DomainError
}

type domainAdvice struct {
}

func NewDomainAdvice() DomainAdvice {
	return &domainAdvice{}
}

func (advice *domainAdvice) TranslateError(err errorInfrastructure.InfrastructureError) DomainError {
	if err == nil {
		return nil
	}
	switch err.GetCode() {
	case errorInfrastructure.ItemNotFound:
		return NewItemNotFoundError(err.GetMessage())
	case errorInfrastructure.SQLError:
		return NewDatabaseError("error while performing database work")
	case errorInfrastructure.ParseFileError:
		return NewConfigurationError("error while parsing configuration file")
	case errorInfrastructure.HashingError:
		return NewCryptoError("error while hashing password")
	default:
		return NewGenericError(err.GetMessage())
	}
}
