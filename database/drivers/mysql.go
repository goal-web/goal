package drivers

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
	"github.com/qbhy/goal/utils"
)

type Mysql struct {
	*sqlx.DB
	prefix string
}

func (this *Mysql) Exec(query string, args ...interface{}) (contracts.Result, error) {
	return this.DB.Exec(query, args...)
}

func MysqlConnector(config contracts.Fields) contracts.DBConnection {
	dsn := utils.GetStringField(config, "unix_socket")
	if dsn == "" {
		dsn = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
			utils.GetStringField(config, "username"),
			utils.GetStringField(config, "password"),
			utils.GetStringField(config, "host"),
			utils.GetStringField(config, "port"),
			utils.GetStringField(config, "database"),
			utils.GetStringField(config, "charset"),
		)
	}
	db, err := sqlx.Connect("mysql", dsn)
	db.SetMaxOpenConns(utils.GetIntField(config, "max_connections"))
	db.SetMaxIdleConns(utils.GetIntField(config, "max_idles"))

	if err != nil {
		logs.WithError(err).WithField("config", config).Fatal("mysql数据库初始化失败")
	}
	return &Mysql{db, utils.GetStringField(config, "prefix")}
}
