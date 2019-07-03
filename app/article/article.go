package article

import (
	"github.com/labstack/echo/v4"
	"goapp/app/article/controllers"
	"goapp/app/article/middleware"
)

func Register(e *echo.Echo) {
	e.GET("articles", controllers.FetchAll)
	e.GET("articles/:name", controllers.FetchOne, middleware.ArticleNameIsValid)
}