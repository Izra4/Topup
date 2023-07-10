package sdk

import (
	"TopUpWEb/entity"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"os"
	"time"
)

func Token(data entity.Admin) (string, error) {
	expStr := os.Getenv("EXP")
	var exp time.Duration
	exp, err := time.ParseDuration(expStr)
	if expStr == "" || err != nil {
		exp = time.Hour / 2
	}
	//generate
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entity.NewAdminClaims(data.ID, exp))
	tokenStr, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", nil
	}
	return tokenStr, nil
}

func DecodeToken(signedToken string, ptrClaims jwt.Claims, KEY string) (string, error) {

	token, err := jwt.ParseWithClaims(signedToken, ptrClaims, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC) // method used to sign the token
		if !ok {
			// wrong signing method
			return "", errors.New("wrong signing method")
		}
		return []byte(KEY), nil
	})

	if err != nil {
		// parse failed
		return "", fmt.Errorf("token has been tampered with")
	}

	if !token.Valid {
		// token is not valid
		return "", fmt.Errorf("invalid token")
	}

	return signedToken, nil
}
