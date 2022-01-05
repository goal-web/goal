package drivers

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/goal/utils"
)

type PostgreSql struct {
	base
}

func PostgreSqlConnector(config contracts.Fields) contracts.DBConnection {
	db, err := sqlx.Connect("postgres", fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		utils.GetStringField(config, "host"),
		utils.GetStringField(config, "port"),
		utils.GetStringField(config, "username"),
		utils.GetStringField(config, "password"),
		utils.GetStringField(config, "database"),
		utils.GetStringField(config, "sslmode"),
	))
	db.SetMaxOpenConns(utils.GetIntField(config, "max_connections"))
	db.SetMaxIdleConns(utils.GetIntField(config, "max_idles"))

	if err != nil {
		logs.WithError(err).WithField("config", config).Fatal("postgreSql 数据库初始化失败")
	}
	return &PostgreSql{base{db}}
}
