package routes

import (
	"goodjobs/controllers/buildings"
	"goodjobs/controllers/complexes"
	"goodjobs/controllers/roles"
	"goodjobs/controllers/units"
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
	UnitController units.UnitController
}

func (ctrl *RouteControllerList) RouteRegister(e *echo.Echo) {

	// user := e.Group("user", middleware.JWTWithConfig(ctrl.JWTMiddleware))

	// user.PUT("/:id", ctrl.UserController.UpdateUserByID)
	jwt := middleware.JWTWithConfig(ctrl.JWTMiddleware)
	
	e.DELETE("/user/:id", ctrl.UserController.DeleteUserByID, jwt)
	e.POST("register", ctrl.UserController.RegisterUser)
	e.POST("login", ctrl.UserController.LoginUser)
	e.GET("user/:id", ctrl.UserController.GetByID)
	e.POST("user", ctrl.UserController.GetByEmail)
	e.GET("users", ctrl.UserController.GetAllUsers)
	e.PUT("user/:id", ctrl.UserController.UpdateUserByID)

	e.POST("role", ctrl.RoleController.Add)
	e.GET("roles", ctrl.RoleController.GetAll)
	e.DELETE("role/:id", ctrl.RoleController.Delete)
	
	e.POST("complex", ctrl.ComplexController.Add)
	e.GET("complexes", ctrl.ComplexController.GetAll)
	e.GET("complex/:id", ctrl.ComplexController.GetByID)
	e.PUT("complex/:id", ctrl.ComplexController.Edit)
	e.DELETE("complex/:id", ctrl.ComplexController.Delete)

	e.POST("building", ctrl.BuildingController.Add)
	e.GET("buildings", ctrl.BuildingController.GetAll)
	e.GET("building/:id", ctrl.BuildingController.GetByID)
	e.PUT("building/:id", ctrl.BuildingController.Edit)
	e.DELETE("building/:id", ctrl.BuildingController.Delete)

	e.POST("unit", ctrl.UnitController.Add)
	e.GET("units", ctrl.UnitController.GetAll)
	e.GET("unit/:id", ctrl.UnitController.GetByID)
	e.PUT("unit/:id", ctrl.UnitController.Edit)
	e.DELETE("unit/:id", ctrl.BuildingController.Delete)
}