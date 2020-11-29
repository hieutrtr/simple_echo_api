package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/echo_api_example/handler"
	mid "github.com/echo_api_example/middleware"
)

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "It's ok bro !")
	})
	mid.UseMiddlewares(e)
	g := e.Group("/users")
	h := handler.New()
	h.RegisterUserHandler(g)

	e.Logger.Fatal(e.Start(":3000"))
}
