package main

import (
	_middleware "goodjobs/app/middlewares"
	"goodjobs/app/routes"
	"goodjobs/driver/mysql"
	"log"
	"time"

	userUseCase "goodjobs/business/users"
	userController "goodjobs/controllers/users"
	userRepo "goodjobs/driver/repository/users"

	roleUseCase "goodjobs/business/roles"
	roleController "goodjobs/controllers/roles"
	roleRepo "goodjobs/driver/repository/roles"

	complexUseCase "goodjobs/business/complexes"
	complexController "goodjobs/controllers/complexes"
	complexRepo "goodjobs/driver/repository/complexes"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

func init() {
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service Run on Debug mode")
	}
}

func DBMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&userRepo.User{})
	DB.AutoMigrate(&roleRepo.Role{})
	DB.AutoMigrate(&complexRepo.Complex{})
}



func main() {
	ConfigDB := mysql.ConfigDB{
		DB_Username: viper.GetString("database.user"),
		DB_Password: viper.GetString("database.pass"),
		DB_Host:     viper.GetString("database.host"),
		DB_Port:     viper.GetString("database.port"),
		DB_Database: viper.GetString("database.name"),
	}

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString("JWT.secretKey"),
		ExpiresDuration: viper.GetInt("JWT.expired_time"),
	}

	DB := ConfigDB.InitialDB()
	DBMigrate(DB)
	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()
	e.Use(middleware.CORS())
	userRepoInterface := userRepo.NewUserRepo(DB)
	userUseCaseInterface := userUseCase.NewUseCase(userRepoInterface, timeoutContext, &_middleware.ConfigJWT{})
	userUseControllerInterface := userController.NewUserController(userUseCaseInterface)

	roleRepoInterface := roleRepo.NewRoleRepo(DB)
	roleUseCaseInterface := roleUseCase.NewUseCase(roleRepoInterface, timeoutContext)
	roleUseControllerInterface := roleController.NewRoleController(roleUseCaseInterface)

	complexRepoInterface := complexRepo.NewComplexRepo(DB)
	complexUseCaseInterface := complexUseCase.NewUseCase(complexRepoInterface, timeoutContext)
	complexUseControllerInterface := complexController.NewComplexController(complexUseCaseInterface)




	routesInit := routes.RouteControllerList{
		UserController: *userUseControllerInterface,
		RoleController: *roleUseControllerInterface,
		ComplexController: *complexUseControllerInterface,
		JWTMiddleware: configJWT.Init(),
	}
	routesInit.RouteRegister(e)
	log.Fatal(e.Start(viper.GetString("server.address")))

}