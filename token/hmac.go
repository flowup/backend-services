package token

import (
	"time"
	"github.com/dgrijalva/jwt-go"
	"fmt"
	"errors"
)

type JWTTokenServiceHMAC struct {
	secret []byte
}


func NewJWTTokenServiceHMAC(secret []byte) *JWTTokenServiceHMAC{
	return &JWTTokenServiceHMAC{
		secret: secret,
	}
}


func (j *JWTTokenServiceHMAC) Create(expiration time.Duration) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"exp": time.Now().Add(expiration).Unix(),
	})

	tokenString, err := token.SignedString(j.secret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func (j *JWTTokenServiceHMAC) Validate(tokenString string) bool {
	t, err := jwt.Parse(tokenString, j.getSecretFunc)
	if err != nil || !t.Valid {
		return false
	}
	return true
}


func (j *JWTTokenServiceHMAC) SetClaimValue(tokenString string, key string, value interface{}) (string, error) {
	token, err := j.Parse(tokenString)
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("Token not valid")
	}
	token.Claims.(jwt.MapClaims)[key] = value

	newString, err := token.SignedString(j.secret)
	if err != nil {
		return "", err
	}

	return newString, nil
}


func (j *JWTTokenServiceHMAC) GetClaimValue(tokenString string, key string) (interface{}, error) {
	token, err := j.Parse(tokenString)
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, errors.New("Token not valid")
	}

	value, found := token.Claims.(jwt.MapClaims)[key]
	if !found {
		return nil, errors.New("Claim not found")
	}
	return value, nil
}


func (j *JWTTokenServiceHMAC) Parse(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, j.getSecretFunc)
}

func(j *JWTTokenServiceHMAC) getSecretFunc (token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("Unexpected signing method, got: %v expected HS256", token.Header["alg"])
	}
	return j.secret, nil
}

