package routes

import (
	"goodjobs/controllers/buildings"
	"goodjobs/controllers/complexes"
	"goodjobs/controllers/roles"
	"goodjobs/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	JWTMiddleware  middleware.JWTConfig
	UserController users.UserController
	RoleController roles.RoleController
	ComplexController complexes.ComplexController
	BuildingController buildings.BuildingController
}

func (ctrl *RouteControllerList) RouteRegister(e *echo.Echo) {

	user := e.Group("user", middleware.JWTWithConfig(ctrl.JWTMiddleware))

	user.PUT("/:id", ctrl.UserController.UpdateUserByID)
	user.DELETE("/:id", ctrl.UserController.DeleteUserByID)
	e.POST("user/register", ctrl.UserController.RegisterUser)
	e.POST("user/login", ctrl.UserController.LoginUser)
	e.GET("user/:id", ctrl.UserController.GetByID)
	e.GET("users", ctrl.UserController.GetAllUsers)
	// e.PUT("user/:id", ctrl.UserController.UpdateUserByID)
	// e.DELETE("user/:id", ctrl.UserController.DeleteUserByID)

	e.POST("role/add", ctrl.RoleController.Add)
	e.GET("roles", ctrl.RoleController.GetAll)
	e.DELETE("role/:id", ctrl.RoleController.Delete)
	
	e.POST("complex/add", ctrl.ComplexController.Add)
	e.GET("complexes", ctrl.ComplexController.GetAll)
	e.DELETE("complex/:id", ctrl.ComplexController.Delete)

	e.POST("building/add", ctrl.BuildingController.Add)
	e.GET("buildings", ctrl.BuildingController.GetAll)
	e.PUT("building/:id", ctrl.BuildingController.Edit)
	e.DELETE("building/:id", ctrl.BuildingController.Delete)
}