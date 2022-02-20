package migrations

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/migrations"
	"github.com/golang-module/carbon/v2"
)

func init() {
	Migrations = append(Migrations, contracts.Migrate{
		CreatedAt:  carbon.Parse("2022-02-15 15:43:39").Carbon2Time(),
		Connection: "sqlite",
		Name:       "2022_02_15_154339_create_users",
		Up: migrations.Exec(`create table users
(
    id   integer     not null
        constraint users_pk
            primary key autoincrement,
    name varchar(20) not null
);
`),
		Down: migrations.Exec("drop table if exists users;"),
	})
}
