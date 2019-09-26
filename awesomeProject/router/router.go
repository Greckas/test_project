package router

import (
	"awesomeProject/handler"

	"github.com/labstack/echo"
)

func Admin(api *echo.Echo, adminHandler *handler.Admin) {
	//ADMIN section
	admin := api.Group("/admin")
	admin.GET("/login", adminHandler.Login)
	admin.POST("/login/go", adminHandler.LoginGo)
	admin.GET("/logout", adminHandler.Logout)
	admin.GET("/home", adminHandler.Home)
	admin.GET("/dashboard", adminHandler.Dashboard)

	//	User

}
