package token

import (
  "time"
  "github.com/dgrijalva/jwt-go"
)

// Service is the interface of tokenService implementation
type Service interface {
  Create(expiration time.Duration) (string, error)
  Parse(tokenString string) (*jwt.Token, error)
  Validate(tokenString string) bool
}
