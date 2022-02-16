package constants

import "errors"

type DeployEnvironment string

const (
	DevEnvironment     DeployEnvironment = "dev"
	QAEnvironment      DeployEnvironment = "qa"
	StgEnvironment     DeployEnvironment = "stg"
	OcuEnvironment     DeployEnvironment = "ocu"
	ProdEnvironment    DeployEnvironment = "prod"
	UnknownEnvironment DeployEnvironment = "unknown"
)

type Environment interface {
	Value() DeployEnvironment
	IsProductiveEnvironment() bool
}

func (env DeployEnvironment) Value() DeployEnvironment {
	return env
}

func (env DeployEnvironment) IsProductiveEnvironment() bool {
	return env == OcuEnvironment || env == ProdEnvironment
}

func NewEnvironment(env string) (DeployEnvironment, error) {
	switch env {
	case string(DevEnvironment):
		return DevEnvironment, nil
	case string(QAEnvironment):
		return QAEnvironment, nil
	case string(StgEnvironment):
		return StgEnvironment, nil
	case string(OcuEnvironment):
		return OcuEnvironment, nil
	case string(ProdEnvironment):
		return ProdEnvironment, nil
	default:
		return UnknownEnvironment, errors.New("unknown environment")
	}
}
