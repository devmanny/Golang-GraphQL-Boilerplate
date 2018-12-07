package thing

import (
	"github.com/labstack/echo"
)

// Router ...
func Router(router *echo.Group) {

	router.GET("/", Index)
	router.POST("/", Index)
	router.PUT("/", Index)
	router.DELETE("/", Index)

}
