package token

import (
  "time"
  "github.com/dgrijalva/jwt-go"
  "crypto/rsa"
  "errors"
  "fmt"
)

// JWTTokenServiceRSA is implementation of TokenService
// using jwt-go library
type JWTTokenServiceRSA struct {
  verifyKey *rsa.PublicKey
  signKey   *rsa.PrivateKey
}

// NewJWTokenServiceRSA is a factory method for JWTTokenServiceRSA
func NewJWTokenServiceRSA(verKey *rsa.PublicKey, signKey *rsa.PrivateKey) *JWTTokenServiceRSA {
  return &JWTTokenServiceRSA{
    verifyKey: verKey,
    signKey:   signKey,
  }
}

// Create will create a token with duration given by parameter
func (j *JWTTokenServiceRSA) Create(expiration time.Duration) (string, error) {
  token := jwt.New(jwt.GetSigningMethod("RS256"))
  token.Claims.(jwt.MapClaims)["exp"] = time.Now().Add(expiration).Unix()

  //sign token
  tokenString, err := token.SignedString(j.signKey)
  if err != nil {
    return "", err
  }

  return tokenString, nil
}

// SetClaimValue will set a value of a claim of token
// all given by parameter if an error occurs it is returned
// if not new token is returned in form of a string
func (j *JWTTokenServiceRSA) SetClaimValue(tokenString string, key string, value interface{}) (string, error) {
  token, err := j.Parse(tokenString)
  if err != nil {
    return "", err
  }

  if !token.Valid {
    return "", errors.New("Token not valid")
  }
  token.Claims.(jwt.MapClaims)[key] = value

  newString, err := token.SignedString(j.signKey)
  if err != nil {
    return "", err
  }

  return newString, nil
}

// GetClaimValue will return a value of a claim
// if an error occurs it is returned
// if not, value of a claim is returned
func (j *JWTTokenServiceRSA) GetClaimValue(tokenString string, key string) (interface{}, error) {
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

// Parse will parse the token from string given
// by parameter and return it
func (j *JWTTokenServiceRSA) Parse(tokenString string) (*jwt.Token, error) {
  return jwt.Parse(tokenString, j.getKeyFunc)
}

// Validate will check if token given by parameter
// is valid
func (j *JWTTokenServiceRSA) Validate(tokenString string) bool {
  reqToken, err := jwt.Parse(tokenString, j.getKeyFunc)
  if err != nil || !reqToken.Valid {
    return false
  }
  return true
}

func (j *JWTTokenServiceRSA) getKeyFunc(token *jwt.Token) (interface{}, error) {
  if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
    return nil, fmt.Errorf("Unexpected signing method, got: %v expected RS256", token.Header["alg"])
  }

  return j.verifyKey, nil
}
