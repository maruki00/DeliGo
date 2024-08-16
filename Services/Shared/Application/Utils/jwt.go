package shared_utils

import (
	shared_configs "delivery/Services/Shared/Application/Configs"
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
)

func JwtToken(email, fullname string) (string, error) {

	conf, err := shared_configs.GetConfig()
	if err != nil {
		return "", errors.New("Config error")
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"email":    email,
			"fullname": fullname,
			"exp":      time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenString, err := token.SignedString(conf.Jwt.Secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
