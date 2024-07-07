package initializers

import (
	"database/sql"
	"log/slog"
	"os"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/millionsmonitoring/millionsgocore/configs"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func InitDB() *bun.DB {
	url := os.Getenv("DB_URL")
	if url == "" {
		slog.Error("DB_URL is not set in the environment")
		panic("DB_URL is not set in the environment")
	}
	dbUrl := strings.Split(url, "//")[1]
	if configs.IsDevelopment() {
		slog.Debug("the db url form the env is ", "url", dbUrl)
	}
	sqlDB, err := sql.Open("mysql", dbUrl)
	if err != nil {
		slog.Error("Error connecting to database while starting application", "error", err)
		panic("error connecting to database")
	}
	db := bun.NewDB(sqlDB, mysqldialect.New())
	if err := db.Ping(); err != nil {
		slog.Error("Error pinging to database while starting application", "error", err)
		panic("error pinging to database")
	}
	return db
}
