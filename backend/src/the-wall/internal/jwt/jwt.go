package jwt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	phoneNumberKey     = "phone_number"
	expiresKey         = "exp"
	expirationDuration = 3 * time.Hour
)

var signingMethod = jwt.SigningMethodHS256

type JWT struct {
	key string
}

func NewJWT(key string) *JWT {
	return &JWT{key: key}
}

func (j *JWT) GetSignedJWT(phoneNumber string) (string, error) {
	claims := jwt.MapClaims{
		phoneNumberKey: phoneNumber,
		expiresKey:     time.Now().Add(expirationDuration).Unix(),
	}

	return jwt.NewWithClaims(signingMethod, claims).SignedString([]byte(j.key))
}

func (j *JWT) CheckValidity(token string) (bool, error) {
	claims := jwt.MapClaims{}
	parse, err := jwt.ParseWithClaims(
		token,
		claims,
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.key), nil
		},
	)
	if err != nil {
		return false, err
	}

	if !parse.Valid {
		return false, nil
	}

	return true, nil
}
