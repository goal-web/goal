package drivers

import (
	"github.com/goal-web/contracts"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/goal/supports/utils"
)

type Sqlite struct {
	base
}

func SqliteConnector(config contracts.Fields) contracts.DBConnection {
	db, err := sqlx.Connect("sqlite3", utils.GetStringField(config, "database"))

	if err != nil {
		logs.WithError(err).WithField("config", config).Fatal("sqlite 数据库初始化失败")
	}

	return &Sqlite{base{db}}
}
