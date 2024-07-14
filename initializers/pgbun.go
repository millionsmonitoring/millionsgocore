package initializers

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/millionsmonitoring/millionsgocore/helpers"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

type PGDBOptions struct {
	DNS string `json:"dns" yaml:"dns"`
}

func (PGDBOptions) DefaultConfig() any {
	return PGDBOptions{
		DNS: "postgres://username:password@host:port/database?sslmode=disable",
	}
}

func InitPGBun(ctx context.Context) *bun.DB {
	options, err := helpers.CheckOrParseConfig[PGDBOptions]("database.yml")
	if err != nil {
		panic(fmt.Sprintf("error in parsing db config: %s", err))
	}
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(options.DNS)))
	if err := sqldb.PingContext(ctx); err != nil {
		panic(fmt.Sprintf("error in connecting to db: %s", err))
	}
	db := bun.NewDB(sqldb, pgdialect.New())
	if err := db.Ping(); err != nil {
		panic(fmt.Sprintf("error pinging to database %s", err))
	}
	return db
}
