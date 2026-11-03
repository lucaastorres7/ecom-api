package auth

import (
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/lucaastorres7/ecom-api/config"
)

func CreateToken(secret string, userId int) (string, error) {
	expiration := time.Second * time.Duration(config.Envs.TokenExpirationInSeconds)

	// Claims to put metadata inside it
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": strconv.Itoa(userId),
		"expiredAt": time.Now().Add(expiration).Unix(),
	})

	secretByte := []byte(secret)
	tokenStr, err := token.SignedString(secretByte)
	if err != nil {
		return "", err
	}

	return tokenStr, nil
}