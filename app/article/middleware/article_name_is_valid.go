package middleware

import (
	"net/http"
	"github.com/labstack/echo/v4"
)

func ArticleNameIsValid(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		name := c.Param("name")
		validNames := map[string]bool {
			"fun": true,
			"times": true,
			"great": true,
		}

		if !validNames[name] {
			return c.String(http.StatusNotFound, "this article doesn't exist, guy")
		}

		return next(c)
	}
}