package thing

import (
	"net/http"

	"github.com/labstack/echo"
)

// Index ...
func Index(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "hello 👋🏻 from things",
	})
}
