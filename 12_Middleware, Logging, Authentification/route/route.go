package route

import (
	"middleware/controller"
	"middleware/middlewares"
	m "middleware/middlewares"

	"github.com/labstack/echo/v4"
	mid "github.com/labstack/echo/v4/middleware"
)

func New() *echo.Echo {
	e := echo.New()
	e.GET("/users", controller.GetUsersController)
	e.POST("/users", controller.CreateUserController)
	m.LogMiddleware(e)
	e.GET("/books", controller.GetBooksController)
	e.POST("/books", controller.CreateBookController)
	

	eAuthBasic : = e.Group"/auth")
	eAuthBasic.Use(mid.BasicAuth(middleware.BasicAuthDB))
	eAuthBasic.Get("/users". controllers.GetUserscontroller)
	return e
}


