package auth

import (
	"api/src/config"
	jwt "github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateToken(userID uint64) (string, error) {
	perms := jwt.MapClaims{}
	perms["authorized"] = true
	perms["exp"] = time.Now().Add(time.Hour * 6).Unix()
	perms["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, perms)

	return token.SignedString([]byte(config.SecretKey))

}
