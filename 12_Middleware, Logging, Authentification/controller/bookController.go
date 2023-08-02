package controller

import (
	"middleware/config"
	"middleware/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetBooksController(c echo.Context) error {
	var books []model.Book

	if err := config.DB.Find(&books).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success get all books",
		"data":    books,
	})
}

func CreateBookController(c echo.Context) error {

	book := model.Book{}
	c.Bind(&book)

	if err := config.DB.Save(&book).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, map[string]interface{}{
		"message": "success create new book",
		"book":    book,
	})
}
