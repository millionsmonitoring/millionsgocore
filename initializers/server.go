package initializers

import (
	"context"
	"log/slog"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"github.com/millionsmonitoring/millionsgocore/configs"
	"github.com/millionsmonitoring/millionsgocore/logger"
	slogecho "github.com/samber/slog-echo"
)

func InitServer() *echo.Echo {
	logger.Info(context.TODO(), "starting the server")
	server := echo.New()
	if configs.IsProduction() {
		server.Logger.SetLevel(log.INFO)
		server.HideBanner = true
		server.HidePort = true
	} else {
		server.Debug = true
		server.Logger.SetLevel(log.DEBUG)
	}
	server.Use(middleware.Recover())
	server.Use(middleware.CORS())
	server.Use(slogecho.New(slog.Default()))
	server.Use(middleware.BodyDump(func(ctx echo.Context, b1, b2 []byte) {
		logger.Info(ctx.Request().Context(), "request", "method", ctx.Request().Method, "uri", ctx.Request().RequestURI, "body", b1)
		logger.Info(ctx.Request().Context(), "response", "method", ctx.Request().Method, "uri", ctx.Request().RequestURI, "body", b2)
	}))
	return server
}
