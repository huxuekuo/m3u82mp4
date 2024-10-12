package utils

import (
	"m3u82mp4/library"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JwtTokenCustom struct {
	Uid int64
	jwt.RegisteredClaims
}

func Token(uid int64) (string, error) {
	mySigningKey := []byte("AllYourBase")
	claims := JwtTokenCustom{
		uid,
		jwt.RegisteredClaims{
			// Also fixed dates can be used for the NumericDate
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24 * 90)),
			Issuer:    "test",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(mySigningKey)

}

func TokenParse(tokenString string) *JwtTokenCustom {
	token, err := jwt.ParseWithClaims(tokenString, &JwtTokenCustom{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil
	})
	if err != nil {
		library.Logger.Sugar().Error(err)
	} else if err == nil {
		if claims, ok := token.Claims.(*JwtTokenCustom); ok {
			return claims
		}
	}
	return nil
}
