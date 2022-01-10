package main

import (
	"bookstore/controller"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.GET("/helloworld", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World")
	})
	e.GET("/view", controller.GetBooks)
	e.POST("/addBook", controller.PostBook)
	e.DELETE("/deleteBook/:id", controller.DeleteBook)
	e.Logger.Fatal(e.Start(":3000"))
}
