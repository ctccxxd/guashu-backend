package jwt

import (
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"jihulab.com/guashu/gs-server/config"
	"jihulab.com/guashu/gs-server/lib/helper"
	"reflect"
)

type TokenData struct {
	UserId   uint64
	ExpireAt int64
}

func ParseToken(tokenString string) (*TokenData, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// hmacSampleSecret is a []byte containing your secret, e.g. []byte("my_secret_key")
		return []byte(config.JwtSecret), nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if ok == false || token.Valid == false {
		return nil, errors.New("validation failed")
	}

	var UserId uint64
	if reflect.TypeOf(claims["user_id"]).String() == "string" {
		UserId = helper.StrToUint64(fmt.Sprintf("%s", claims["user_id"]))
	} else {
		UserId = uint64(claims["user_id"].(float64))
	}

	return &TokenData{
		UserId:   UserId,
		ExpireAt: int64(claims["exp"].(float64)),
	}, nil

}

func MakeToken(data TokenData) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":   data.UserId,
		"expire_at": data.ExpireAt,
	})

	tokenString, _ := token.SignedString([]byte(config.JwtSecret))
	return tokenString
}
