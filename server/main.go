package main

import (
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
	e.GET("/command", runCommand)
}

// e.GET("/users/:id", getUser)
func runCommand(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("run")
	strings.Split(id, "")
	return c.String(http.StatusOK, id)
}
