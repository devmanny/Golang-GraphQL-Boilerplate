package thing

import (
	"github.com/labstack/echo"
)

// Router ...
func Router(router *echo.Group) {

	router.GET("/", GraphQL)
	router.POST("/", GraphQL)

}
