package picture

import (
	"github.com/labstack/echo/v4"
	"goapp/app/picture/controllers"
)

func Register(e *echo.Echo) {
	e.GET("picture", controllers.Show)
}

