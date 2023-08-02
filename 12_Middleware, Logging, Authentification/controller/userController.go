package controller

import (
	"middleware/config"
	"middleware/model"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUserController(c echo.Context) error {
	var users []model.Users

	err := config.DB.Find(&users).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	// get Id from JWT
	// myUserId := middleware.ExtractTokenUserId(c)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    users,
	})
}

func CreateUserController(c echo.Context) error {

	user := model.Users{}
	c.Bind(&user)

	err := config.DB.Save(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": err.Error(),
		})
	}

	token, err := middleware.CreateToken(int(user.ID), user.Name)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "Fail create JWT TOken",
			"status":  err.Error(),
		})
	}

	user.Token = token

	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create user",
		"user":    user,
	})
}
