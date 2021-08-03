package routers

import (
	apiController "crm-golang/controllers/api"

	"github.com/labstack/echo/v4"
)

func SetupRouter() *echo.Echo {
	e := echo.New()

	e.GET("/users", apiController.UserList)

	e.POST("user", apiController.CreateUser)

	return e
}
