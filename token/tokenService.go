package token

import (
  "time"
  "github.com/dgrijalva/jwt-go"
)

type TokenService interface {
  Create(expiration time.Duration) (string, error)
  Parse(tokenString string) (*jwt.Token, error)
  Validate(tokenString string) bool
}
