package user

import (
	"github.com/labstack/echo"
)

// Router ...
func Router(router *echo.Group) {

	router.GET("/", GraphQL)
	router.POST("/", GraphQL)

	// router.GET("/", Index)
	// router.PUT("/", Index)
	// router.DELETE("/", Index)

}
