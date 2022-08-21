package helper

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"math/rand"
	"op-panel/define"
	"time"
)

func RandomString(n int) string {
	s := "0123456789AaBbCcDdEeFfGgHhIiJjKkLlMmNnOoPpQqRrSsTtUuVvWwXxYyZz"
	rand.Seed(time.Now().UnixNano())
	ans := make([]byte, 0, n)
	for i := 0; i < n; i++ {
		ans = append(ans, s[rand.Intn(len(s))])
	}
	return string(ans)
}

func GenerateToken() (string, error) {
	tokenStruct := jwt.NewWithClaims(jwt.SigningMethodHS256, &define.UserClaim{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: &jwt.NumericDate{
				Time: time.Now().Add(time.Hour * 24),
			},
		},
	})
	token, err := tokenStruct.SignedString(define.Key)
	if err != nil {
		return "", err
	}
	return token, nil
}

func ParseToken(token string) error {
	claims, err := jwt.ParseWithClaims(token, &define.UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return define.Key, nil
	})
	if err != nil {
		return err
	}
	if !claims.Valid {
		return errors.New("error Token")
	}
	return nil
}
