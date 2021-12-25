package middlewares

import (
	// controller "goodjobs/controllers"
	// "net/http"

	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	// "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
    ID          uint            /*`json:"id"`*/
    Email     	string          /*`json:"email"`*/
	Name      	string          /*`json:"name"`*/
	Phone		string 			/*`json:"phone"`*/
    Role_ID     uint            /*`json:"role_id"`*/
    jwt.StandardClaims
}

type ConfigJWT struct {
    SecretJWT string
    ExpiresDuration int
}

func (jwtConf *ConfigJWT) Init() middleware.JWTConfig{
    fmt.Println(jwtConf.SecretJWT)
	return middleware.JWTConfig{
		Claims: &JWTCustomClaims{},
		SigningKey: []byte(jwtConf.SecretJWT),
		// ErrorHandlerWithContext: middleware.JWTErrorHandlerWithContext(func(e error, c echo.Context) error {
		// 	return controller.NewErrorResponse(c, http.StatusForbidden, e)
		// }),
	}
}



func (jwtConf *ConfigJWT) GenerateTokenJWT(id uint, email string, name string, phone string, role_id uint) (string, error) {
    claims := JWTCustomClaims{
        id,
        email,
        name,
        phone,
        role_id,
        jwt.StandardClaims{
            ExpiresAt: time.Now().Add(time.Hour*3).Local().Add(time.Hour * time.Duration(int64(jwtConf.ExpiresDuration))).Unix(),
        },
    }
    fmt.Println(claims)
    fmt.Println(jwtConf.SecretJWT)
    t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    token, err := t.SignedString([]byte(jwtConf.SecretJWT))
    return token, err
}
