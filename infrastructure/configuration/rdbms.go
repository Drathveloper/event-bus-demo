package configuration

import (
	"database/sql"
	infrastructure "event-bus-demo/infrastructure/error"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

func BuildDatabase(configuration RdbmsConfiguration) (*sql.DB, infrastructure.InfrastructureError) {
	connectionUrl := fmt.Sprintf("%s://%s:%s@%s:%s/%s?sslmode=disable",
		*configuration.Driver,
		*configuration.User,
		*configuration.Password,
		*configuration.Host,
		*configuration.Port,
		*configuration.Database)
	db, err := sql.Open(*configuration.Driver, connectionUrl)
	if err != nil {
		return nil, infrastructure.NewSQLError(err.Error())
	}
	maxConnectionLifetime, err := time.ParseDuration(*configuration.Pool.MaxConnectionLifetime)
	if err != nil {
		return nil, infrastructure.NewSQLError(err.Error())
	}
	db.SetMaxOpenConns(*configuration.Pool.MaxOpenConnections)
	db.SetMaxIdleConns(*configuration.Pool.MaxIdleConnections)
	db.SetConnMaxLifetime(maxConnectionLifetime)
	return db, nil
}
