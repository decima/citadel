package security

import (
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

type JWTManagerInterface[T any] interface {
	Generate(data T) (string, error)
	Decode(key string) (*T, error)
}

type JWTManager[T any] struct {
	signature []byte
	validity  time.Duration
}

func NewJWTManager[T any](signature []byte, validity time.Duration) *JWTManager[T] {
	return &JWTManager[T]{
		signature: signature,
		validity:  validity,
	}
}
func NewUnlimitedJWTManager[T any](signature []byte) *JWTManager[T] {
	return &JWTManager[T]{
		signature: signature,
		validity:  0,
	}
}

func (m *JWTManager[T]) WithSignature(signature []byte) *JWTManager[T] {
	m.signature = signature
	return m
}
func (m *JWTManager[T]) WithValidity(validity time.Duration) *JWTManager[T] {
	m.validity = validity
	return m
}

func (m *JWTManager[T]) Generate(data T) (string, error) {

	mapClaims := jwt.MapClaims{
		"iat":  time.Now().Unix(),
		"data": data,
	}

	if m.validity > 0 {
		mapClaims["exp"] = time.Now().Add(m.validity).Unix()
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, mapClaims)
	return token.SignedString(m.signature)

}

func (m *JWTManager[T]) Decode(tokenString string) (*T, error) {
	// Parse the token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Ensure that the token's signing method is what you expect
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			return nil, errors.New("invalid claims")
		}
		if iat, isSet := claims["iat"]; isSet {
			if time.Now().Unix() < int64(iat.(float64)) {
				return nil, errors.New("token not yet valid")
			}
		}
		if exp, isSet := claims["exp"]; isSet {
			if time.Now().Unix() > int64(exp.(float64)) {
				return nil, errors.New("token expired")
			}
		}

		return m.signature, nil
	})
	claims, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}

	jbyte, err := json.Marshal(claims["data"])
	if err != nil {
		return nil, err
	}

	var data T
	err = json.Unmarshal(jbyte, &data)
	return &data, err
}
