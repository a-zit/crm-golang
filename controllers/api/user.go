package apicontroller

import (
	us "crm-golang/services/api"
	"crm-golang/validator/user"
	"fmt"

	"github.com/labstack/echo/v4"
)

func UserList(c echo.Context) error {
	var userService us.UserService

	users, err := userService.UserList()

	if err != nil {
		c.JSON(500, "Internal Server Error")
	}

	return c.JSON(200, users)
}

func CreateUser(c echo.Context) error {
	var userService us.UserService
	var userRequest user.CreateUserRequest
	// userRequest := new(user.CreateUserRequest)
	if err := c.Bind(&userRequest); err != nil {
		fmt.Println(err)
		return err
	}

	user, err := userService.CreateUser(userRequest)
	// fmt.Println(user)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, "Internal Server Error")
	}

	return c.JSON(200, user)

}
