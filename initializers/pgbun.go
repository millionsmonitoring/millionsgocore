package initializers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"gopkg.in/yaml.v2"
)

type DBConfig interface {
	defaultConfig() any
}

type PGDBOptions struct {
	DNS string `json:"dns" yaml:"dns"`
}

func (PGDBOptions) defaultConfig() any {
	return PGDBOptions{
		DNS: "postgres://postgres:@localhost:5432/test?sslmode=disable",
	}
}

var _ DBConfig = (*PGDBOptions)(nil)

func InitPGBun(ctx context.Context) *bun.DB {
	options, err := checkDatabaseConfigPresence[PGDBOptions]()
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

// Check if database.yml exists, if not create it with default config
func checkDatabaseConfigPresence[T DBConfig]() (T, error) {
	var options T
	configPath := filepath.Join("configs", "database.yml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// File doesn't exist, create it with default config

		defaultConfig := options.defaultConfig()

		data, err := yaml.Marshal(&defaultConfig)
		if err != nil {
			return options, err
		}

		err = os.MkdirAll(filepath.Dir(configPath), 0755)
		if err != nil {
			return options, err
		}

		err = os.WriteFile(configPath, data, 0644)
		if err != nil {
			return options, err
		}

		return options, errors.New("database.yml created in configs folder. Please update with your database details")
	}

	data, err := os.ReadFile(configPath)
	if err != nil {
		return options, err
	}

	err = yaml.Unmarshal(data, &options)
	if err != nil {
		return options, err
	}
	return options, nil
}
