package routes

import (
	"goodjobs/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
}

func (ctrl *RouteControllerList) RouteRegister(e *echo.Echo) {
	// laundro := e.Group("laundro", middleware.JWTWithConfig(ctrlList.JWTMiddleware))
	// laundro.GET("/find-ip", ctrlList.LaundromatController.GetByIP)
	// laundro.GET("/find-name/:name", ctrlList.LaundromatController.GetByName)
	// laundro.GET("/find-category/:categoryId", ctrlList.ProductController.GetLaundromatByCategory)

	user := e.Group("user", middleware.JWTWithConfig(ctrl.JWTMiddleware))

	user.PUT("/:id", ctrl.UserController.UpdateUserByID)
	user.DELETE("/:id", ctrl.UserController.DeleteUserByID)
	e.POST("user/register", ctrl.UserController.RegisterUser)
	e.POST("user/login", ctrl.UserController.LoginUser)
	e.GET("user/:id", ctrl.UserController.GetByID)
	e.GET("users", ctrl.UserController.GetAllUsers)
	// e.PUT("user/:id", ctrl.UserController.UpdateUserByID)
	// e.DELETE("user/:id", ctrl.UserController.DeleteUserByID)

}