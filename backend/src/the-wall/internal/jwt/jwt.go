package jwt

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	phoneNumberKey     = "phone_number"
	expiresKey         = "exp"
	expirationDuration = 3 * time.Hour
	authTokenType      = "Bearer"
)

var (
	signingMethod           = jwt.SigningMethodHS256
	invalidTokenFormatError = errors.New("invalid token format")
)

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

	signedString, err := jwt.NewWithClaims(signingMethod, claims).SignedString([]byte(j.key))
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s %s", authTokenType, signedString), nil
}

func (j *JWT) CheckValidity(token string) (bool, error) {
	tokenSplits := strings.Split(token, " ")
	if len(tokenSplits) != 2 {
		return false, invalidTokenFormatError
	}

	jwtToken, tokenType := tokenSplits[1], tokenSplits[0]
	if tokenType != authTokenType {
		return false, invalidTokenFormatError
	}

	claims := jwt.MapClaims{}
	parse, err := jwt.ParseWithClaims(
		jwtToken,
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
