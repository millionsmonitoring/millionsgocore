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

type PGDBOptions struct {
	DNS string `json:"dns" yaml:"dns"`
}

func InitPGBun(ctx context.Context) *bun.DB {
	options, err := checkDatabaseConfigPresence()
	if err != nil {
		panic(fmt.Sprintf("error in parsing db config: %s", err))
	}
	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(options.DNS)))
	if err := sqldb.PingContext(ctx); err != nil {
		panic(fmt.Sprintf("error in connecting to db: %s", err))
	}
	return bun.NewDB(sqldb, pgdialect.New())
}

// Check if database.yml exists, if not create it with default config
func checkDatabaseConfigPresence() (PGDBOptions, error) {
	var options PGDBOptions
	configPath := filepath.Join("configs", "database.yml")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		// File doesn't exist, create it with default config
		defaultConfig := defaultConfig()

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

func defaultConfig() PGDBOptions {
	return PGDBOptions{
		DNS: "postgres://postgres:@localhost:5432/test?sslmode=disable",
	}
}
