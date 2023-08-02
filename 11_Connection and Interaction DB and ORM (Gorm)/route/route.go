package route

import (
	"gorm/controller"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controller.GetUsersController)
	e.POST("/users", controller.CreateUserController)

	e.GET("/books", controller.GetBooksController)
	e.POST("/books", controller.CreateBookController)

	return e
}
