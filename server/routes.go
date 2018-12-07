package server

import (
	"net/http"

	"app.onca.api/server/api/thing"
	"app.onca.api/server/api/user"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

// ConfigureRoutes ...
func ConfigureRoutes() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiV1 := e.Group("/api/v1") // se puede omitir la versión

	user.Router(apiV1.Group("/users"))
	thing.Router(apiV1.Group("/things"))

	e.GET("/*", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]string{
			"message": "hello 👋🏻",
		})
	})

	http.Handle("/", e)

	return e
}
