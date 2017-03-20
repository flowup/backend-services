package token

import (
  "time"
  "github.com/dgrijalva/jwt-go"
  "crypto/rsa"
  "errors"
  "fmt"
)

// RSAService is implementation of TokenService
// using jwt-go library
type RSAService struct {
  verifyKey *rsa.PublicKey
  signKey   *rsa.PrivateKey
}

// NewRSAService is a factory method for RSAService
func NewRSAService(verKey *rsa.PublicKey, signKey *rsa.PrivateKey) *RSAService {
  return &RSAService{
    verifyKey: verKey,
    signKey:   signKey,
  }
}

// Create will create a token with duration given by parameter
func (j *RSAService) Create(expiration time.Duration) (string, error) {
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
func (j *RSAService) SetClaimValue(tokenString string, key string, value interface{}) (string, error) {
  token, err := j.parse(tokenString)
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
func (j *RSAService) GetClaimValue(tokenString string, key string) (interface{}, error) {
  token, err := j.parse(tokenString)
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
func (j *RSAService) parse(tokenString string) (*jwt.Token, error) {
  return jwt.Parse(tokenString, j.getKeyFunc)
}

// Validate will check if token given by parameter
// is valid
func (j *RSAService) Validate(tokenString string) bool {
  reqToken, err := jwt.Parse(tokenString, j.getKeyFunc)
  if err != nil || !reqToken.Valid {
    return false
  }
  return true
}

func (j *RSAService) getKeyFunc(token *jwt.Token) (interface{}, error) {
  if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
    return nil, fmt.Errorf("Unexpected signing method, got: %v expected RS256", token.Header["alg"])
  }

  return j.verifyKey, nil
}
