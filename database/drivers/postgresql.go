package drivers

import (
	"fmt"
	"github.com/goal-web/application"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/logs"
	"github.com/goal-web/supports/utils"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type PostgresSQL struct {
	*Base
}

func PostgresSqlConnector(config contracts.Fields) contracts.DBConnection {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		utils.GetStringField(config, "host"),
		utils.GetStringField(config, "port"),
		utils.GetStringField(config, "username"),
		utils.GetStringField(config, "password"),
		utils.GetStringField(config, "database"),
		utils.GetStringField(config, "sslmode"),
	))
	if err != nil {
		logs.WithError(err).WithField("config", config).Fatal("postgreSql 数据库初始化失败")
	}
	db.SetMaxOpenConns(utils.GetIntField(config, "max_connections"))
	db.SetMaxIdleConns(utils.GetIntField(config, "max_idles"))

	return &PostgresSQL{&Base{db, application.Get("events").(contracts.EventDispatcher)}}
}
