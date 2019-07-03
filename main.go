package main

import (
	"github.com/labstack/echo/v4"

	"goapp/app/article"
	"goapp/app/picture"
)

func main() {
	e := echo.New()

	registerServices(e)

	e.Logger.Fatal(e.Start(":8080"))
}

func registerServices(e *echo.Echo) {
	article.Register(e)
	picture.Register(e)
}