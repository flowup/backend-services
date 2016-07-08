package token

import (
  "time"
  "github.com/dgrijalva/jwt-go"
  "crypto/rsa"
)

type JWTokenService struct {
  VerifyKey *rsa.PublicKey
  SignKey *rsa.PrivateKey
}

func NewJWTokenService (verKey *rsa.PublicKey, signKey *rsa.PrivateKey) *JWTokenService {
  return &JWTokenService{
    VerifyKey:verKey,
    SignKey:signKey,
  }
}

func (j *JWTokenService) Create(expiration time.Duration) (string, *jwt.Token, error) {
  token := jwt.New(jwt.GetSigningMethod("RS256"))
  token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(expiration).Unix()

  //sign token
  tokenString, err := token.SignedString(j.SignKey)
  if err != nil {
    return "", nil, err
  }

  return tokenString, token, nil
}

func (j *JWTokenService) Parse(tokenString string) (*jwt.Token, error) {
  return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return j.VerifyKey, nil
  })
}

func (j *JWTokenService) Validate(tokenString string) bool {
  reqToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
    return j.VerifyKey, nil
  })
  if err != nil || !reqToken.Valid {
    return false
  }
  return true
}