package initializers

import (
	"context"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/millionsmonitoring/millionsgocore/helpers"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

type MySqlDBOptions struct {
	URL string `json:"url" yaml:"url"`
}

func (MySqlDBOptions) DefaultConfig() any {
	return MySqlDBOptions{
		URL: "mysql://username:password@tcp(host:$port)/database?parseTime=true",
	}
}

func InitMysqlBun(ctx context.Context) *bun.DB {
	options, err := helpers.CheckOrParseConfig[MySqlDBOptions]("database.yml")
	if err != nil {
		panic(fmt.Sprintf("error in parsing db config: %s", err))
	}
	sqlDB, err := sql.Open("mysql", options.URL)
	if err != nil {
		panic(fmt.Sprintf("error in connecting to db: %s", err))
	}
	db := bun.NewDB(sqlDB, mysqldialect.New())
	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("error pinging to database %s", err))
	}
	return db
}
