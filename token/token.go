package token

import (
  "time"
)

// Service is the interface of tokenService implementation
type Service interface {
  Create(expiration time.Duration) (string, error)
  Validate(tokenString string) bool
  SetClaimValue(tokenString string, key string, value interface{}) (string, error)
  GetClaimValue(tokenString string, key string) (interface{}, error)
}
