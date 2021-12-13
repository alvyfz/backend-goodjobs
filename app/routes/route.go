package routes

type RouteControllerList struct {
	JWTMiddleware  middlewares.ConfigJWT
	UserController users.UserController
}

func (ctrl *RouteControllerList) RouteRegister(e *echo.Echo) {
	e.POST("user/register", ctrl.UserController.RegisterUser)
	e.POST("user/login", ctrl.UserController.LoginUser)
	e.GET("user/:id", ctrl.UserController.GetByID)
	e.GET("users", ctrl.UserController.GetAllUsers)
	e.PUT("user/:id", ctrl.UserController.UpdateUserByID)
	e.DELETE("user/:id", ctrl.UserController.DeleteUserByID)

}