package jwt

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

var (
	signingMethod   = jwt.SigningMethodHS256
	ErrInvalidToken = errors.New("invalid token")
)

type JWT struct {
	key                string
	expirationDuration time.Duration
}

func NewJWT(key string, options ...Option) *JWT {
	j := &JWT{
		key:                key,
		expirationDuration: 3 * time.Hour,
	}
	for _, option := range options {
		option(j)
	}
	return j
}

type Option func(*JWT)

func WithExpireDuration(d time.Duration) Option {
	return func(j *JWT) {
		j.expirationDuration = d
	}
}

func (j *JWT) GetSignedJWT(payload Payload) (string, error) {
	jti, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	claims := Claims{
		Payload: payload,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(j.expirationDuration)),
			ID:        jti.String(),
		},
	}

	signedString, err := jwt.NewWithClaims(signingMethod, claims).SignedString([]byte(j.key))
	if err != nil {
		return "", err
	}

	return signedString, nil
}

func (j *JWT) CheckValidity(token string) error {
	parsed, err := jwt.ParseWithClaims(
		token,
		&Claims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.key), nil
		},
		jwt.WithExpirationRequired(),
	)
	if err != nil {
		return errors.Join(ErrInvalidToken, err)
	}

	if !parsed.Valid {
		return ErrInvalidToken
	}

	return nil
}

// ExtractPayload parses the token but does not verify the signature.
func (j *JWT) ExtractPayload(token string) (Payload, error) {
	parsed, _, err := jwt.NewParser().ParseUnverified(token, &Claims{})
	if err != nil {
		return Payload{}, err
	}
	if claims, ok := parsed.Claims.(*Claims); ok {
		return claims.Payload, nil
	}
	return Payload{}, nil
}
