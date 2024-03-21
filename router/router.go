package router

import (
	"database/sql"
	"net/http"

	"github.com/labstack/echo/v4"
)

func SetupRouter(config *Config, db *sql.DB) *echo.Echo {
	return newEcho(config, db)
}

func newEcho(config *Config, db *sql.DB) *echo.Echo {
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// ミドルウェア設定
	/*e.Use(middlewares.ServerVersion(config.Version))
	e.Use(middlewares.RequestID())

	e.Use(middlewares.Recovery(logger))

	e.Use(extension.Wrap(repo, cm))
	e.Use(middlewares.RequestCounter())*/

	api := e.Group("/api")
	api.GET("/ping", func(c echo.Context) error { return c.String(http.StatusOK, http.StatusText(http.StatusOK)) })

	return e
}
