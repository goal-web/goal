package database

import (
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/database/drivers"
)

type ServiceProvider struct {
}

func (this *ServiceProvider) Register(application contracts.Application) {
	application.Singleton("db.factory", func(config contracts.Config) contracts.DBFactory {
		return &Factory{
			config:      config,
			connections: make(map[string]contracts.DBConnection),
			drivers: map[string]contracts.DBConnector{
				"mysql":    drivers.MysqlConnector,
				"postgres": drivers.PostgreSqlConnector,
				"sqlite":   drivers.SqliteConnector,
			},
		}
	})
	application.Singleton("db", func(config contracts.Config, factory contracts.DBFactory) contracts.DBConnection {
		return factory.Connection(config.Get("database").(Config).Default)
	})
}

func (this *ServiceProvider) Start() error {
	return nil
}

func (this *ServiceProvider) Stop() {
}
