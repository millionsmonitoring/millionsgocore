package millionsgocore

// package core

// import (
// 	"fmt"
// 	"log/slog"
// 	"sync"

// 	"github.com/labstack/echo/v4"
// 	"github.com/millionsmonitoring/millionsgocore/config"
// 	"github.com/millionsmonitoring/millionsgocore/core/initializers"
// 	"github.com/millionsmonitoring/millionsgocore/lib/utils/asynqwrapper"
// 	"github.com/uptrace/bun"
// 	"go.mongodb.org/mongo-driver/mongo"
// )

// type application struct {
// 	logger  *slog.Logger
// 	server  *echo.Echo
// 	db      *bun.DB
// 	worker  *asynqwrapper.TaskClient
// 	mongodb *mongo.Client
// }

// var (
// 	app     application
// 	appOnce sync.Once
// )

// func (app *application) init() {
// 	configLogger()
// 	app.logger = initializers.InitLogger()
// 	app.db = initializers.InitDB()
// 	app.server = initializers.InitServer()
// 	mongo, err := initializers.InitMongoDB()
// 	if err != nil {
// 		panic(fmt.Sprintf("unable to start mongodb %+v", err))
// 	}
// 	app.mongodb = mongo
// 	worker, err := initializers.InitAsynq()
// 	if err != nil {
// 		panic(fmt.Sprintf("unable to start asynq worker %+v", err))
// 	}
// 	app.worker = worker
// }

// func App() application {
// 	appOnce.Do(func() {
// 		app.init()
// 	})
// 	return app
// }

// func Script() application {
// 	appOnce.Do(func() {
// 		app.init()
// 	})
// 	return app
// }

// func Test() application {
// 	appOnce.Do(func() {
// 		app.init()
// 	})
// 	return app
// }

// func AsynqWorker() application {
// 	appOnce.Do(func() {
// 		app.init()
// 	})
// 	return app
// }

// func Server() *echo.Echo {
// 	return App().server
// }

// func DB() *bun.DB {
// 	return App().db
// }

// func L() *slog.Logger {
// 	return App().logger
// }

// func TaskClient() *asynqwrapper.TaskClient {
// 	return App().worker
// }

// func MongoDB() *mongo.Client {
// 	return App().mongodb
// }

// func configLogger() {
// 	slog.Info("The app env is ", "env", config.Env())
// 	if config.IsDevelopment() {
// 		slog.Info("The app is running in development mode")
// 		slog.Info("The app", "settings", config.Settings())
// 		slog.Info("The app", "secrets", config.Secrets())
// 	}
// }
