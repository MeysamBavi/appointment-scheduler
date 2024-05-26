package jwt

import (
	"github.com/golang-jwt/jwt/v5"
)

type Payload struct {
	UserId      uint   `json:"user_id,omitempty"`
	PhoneNumber string `json:"phone_number,omitempty"`
}

type Claims struct {
	jwt.RegisteredClaims
	Payload
}
