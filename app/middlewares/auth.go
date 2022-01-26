package middlewares

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4/middleware"
)

type JWTCustomClaims struct {
    ID          uint       
    Email     	string        
	Name      	string       
	Phone		string 	
    Role_ID     uint            
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
