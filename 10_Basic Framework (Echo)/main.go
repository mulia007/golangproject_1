package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

var users []User

// -------------------- controller --------------------

// get all users
func GetUsersController(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success get all users",
		"users":    users,
	})
}

// get user by id
func GetUserController(c echo.Context) error {
	idParam := c.Param("id")
	id := parseIdParam(idParam)

	for _, user := range users {
		if user.Id == id {
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success get user by id",
				"user":     user,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// Helper function to parse id parameter to int
func parseIdParam(idParam string) int {
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return 0
	}
	return id
}

// delete user by id
func DeleteUserController(c echo.Context) error {
	idParam := c.Param("id")
	id := parseIdParam(idParam)

	for i, user := range users {
		if user.Id == id {
			// Delete the user from the slice
			users = append(users[:i], users[i+1:]...)
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success delete user by id",
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// update user by id
func UpdateUserController(c echo.Context) error {
	idParam := c.Param("id")
	id := parseIdParam(idParam)

	var updatedUser User
	// binding data
	if err := c.Bind(&updatedUser); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"messages": "invalid request data",
		})
	}

	for i, user := range users {
		if user.Id == id {
			// Update the user in the slice
			updatedUser.Id = id
			users[i] = updatedUser
			return c.JSON(http.StatusOK, map[string]interface{}{
				"messages": "success update user by id",
				"user":     updatedUser,
			})
		}
	}

	return c.JSON(http.StatusNotFound, map[string]interface{}{
		"messages": "user not found",
	})
}

// create new users
func CreateUserController(c echo.Context) error {
	// binding data
	user := User{}
	c.Bind(&user)

	if len(users) == 0 {
		user.Id = 1
	} else {
		newId := users[len(users)-1].Id + 1
		user.Id = newId
	}
	users = append(users, user)
	return c.JSON(http.StatusOK, map[string]interface{}{
		"messages": "success create user",
		"user":     user,
	})
}

// ---------------------------------------------------
func main() {
	e := echo.New()
	// routing with query parameter
	e.GET("/users", GetUsersController)
	e.POST("/users", CreateUserController)

	// start the server, and log if it fails
	e.Logger.Fatal(e.Start(":8000"))
}
