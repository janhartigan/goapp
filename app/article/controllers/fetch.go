package controllers

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func FetchAll(c echo.Context) error {
	return c.String(http.StatusOK, "It's a list of articles")
}

func FetchOne(c echo.Context) error {
	name := c.Param("name")

	return c.String(http.StatusOK, "You're reading this article about " + name)
}