package database

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/goal/database/drivers"
)

type ServiceProvider struct {
}

func (this *ServiceProvider) Register(application contracts.Application) {
	application.Singleton("db.factory", func(config contracts.Config) contracts.DBFactory {
		return &Factory{
			config:      config,
			dbConfig:    config.Get("database").(Config),
			connections: make(map[string]contracts.DBConnection),
			drivers: map[string]contracts.DBConnector{
				"mysql":    drivers.MysqlConnector,
				"postgres": drivers.PostgresSqlConnector,
				"sqlite":   drivers.SqliteConnector,
			},
		}
	})
	application.Singleton("db", func(config contracts.Config, factory contracts.DBFactory) contracts.DBConnection {
		return factory.Connection()
	})
}

func (this *ServiceProvider) Start() error {
	return nil
}

func (this *ServiceProvider) Stop() {
}
