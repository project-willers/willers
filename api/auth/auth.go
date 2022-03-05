package auth

import (
	"crypto/sha512"
	"encoding/hex"

	"github.com/golang-jwt/jwt"

	"willers-api/model"
)

type Token struct {
	Jwt string `json:"jwt"`
}

var hmacSecret = []byte("SECRETKEY")

func CreateToken(user *model.Account) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name":     user.Name,
		"email":    user.Email,
		"password": user.Password,
		// "exp":      time.Now().Add(time.Minute * 1).Unix(),
	})

	tokenString, err := token.SignedString(hmacSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func HashStr(trg string) string {
	sha512 := sha512.Sum512([]byte(trg))
	hashed := hex.EncodeToString(sha512[:])
	return hashed
}
