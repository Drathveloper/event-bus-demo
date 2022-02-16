package configuration

type ApplicationConfiguration struct {
	Gin   *GinConfiguration   `mapstructure:"gin" validate:"required"`
	Event *EventConfiguration `mapstructure:"event" validate:"required"`
	Rdbms *RdbmsConfiguration `mapstructure:"rdbms" validate:"required"`
}

type EventConfiguration struct {
	ChannelBufferSize *int `mapstructure:"channel-buffer-size" validate:"required,min=1"`
	MaxWorkers        *int `mapstructure:"max-workers" validate:"required"`
}

type GinConfiguration struct {
	Environment *string `mapstructure:"environment" validate:"required,oneof=dev qa stg ocu prod"`
	Port        *int    `mapstructure:"port" validate:"required"`
}

type RdbmsConfiguration struct {
	Driver   *string                 `mapstructure:"driver" validate:"required"`
	Host     *string                 `mapstructure:"host" validate:"required"`
	Port     *string                 `mapstructure:"port" validate:"required"`
	Database *string                 `mapstructure:"database" validate:"required"`
	User     *string                 `mapstructure:"user" validate:"required"`
	Password *string                 `mapstructure:"password" validate:"required"`
	Pool     *RdbmsPoolConfiguration `mapstructure:"pool" validate:"required"`
}

type RdbmsPoolConfiguration struct {
	MaxIdleConnections    *int    `mapstructure:"max-idle-connections" validate:"required"`
	MaxOpenConnections    *int    `mapstructure:"max-open-connections" validate:"required"`
	MaxConnectionLifetime *string `mapstructure:"max-connection-lifetime" validate:"required"`
}
