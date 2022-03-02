package migrations

import (
	"github.com/goal-web/contracts"
	"github.com/goal-web/database/migrations"
	"github.com/golang-module/carbon/v2"
)

func init() {
	Migrations = append(Migrations, contracts.Migrate{
		CreatedAt:  carbon.Parse("2022-03-02 21:22:39").Carbon2Time(),
		Connection: "mysql",
		Name:       "2022_02_15_154339_create_users",
		Up: migrations.Exec(`
create table failed_jobs
(
    id         integer not null
        constraint failed_jobs_pk
            primary key autoincrement,
    payload    text    not null,
    exception  text    not null,
    queue      text    not null,
    connection text    not null
);
`),
		Down: migrations.Exec("drop table if exists failed_jobs;"),
	})
}
