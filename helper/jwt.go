package helper

import (
	"fmt"
	"time"

	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/config"
	jwt "github.com/golang-jwt/jwt/v4"
)

type JWTClaim struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func GenerateJWT(email string) (tokenStr string, err error) {

	expTime := time.Now().Add(time.Hour * time.Duration(config.Config.Jwt.ExpDurHour))
	claims := &JWTClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	tokenStr, err = token.SignedString([]byte(config.Config.Jwt.SecretKey))

	return
}

func ValidateToken(signedToken string) (JWTClaim, error) {

	claim := &JWTClaim{}

	token, err := jwt.ParseWithClaims(signedToken, claim, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("there was an Error ")
		}
		return []byte(config.Config.Jwt.SecretKey), nil
	})

	if err != nil || !token.Valid {
		return *claim, err
	}

	if claim.ExpiresAt < time.Now().Local().Unix() {
		return *claim, jwt.ErrTokenExpired
	}

	return *claim, nil
}
