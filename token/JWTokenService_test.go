package token

import (
  "github.com/stretchr/testify/suite"
  "testing"
  "io/ioutil"
  "github.com/dgrijalva/jwt-go"
  "crypto/rsa"
  "time"
  "github.com/stretchr/testify/assert"
)

func initRSAKeys() (*rsa.PublicKey, *rsa.PrivateKey, error) {
  signBytes, err := ioutil.ReadFile("../certificates/app.rsa")
  if err != nil {
    return err, nil, nil
  }
  signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
  if err != nil {
    return err, nil, nil
  }
  verifyBytes, err := ioutil.ReadFile("../certificates/app.rsa.pub")
  if err != nil {
    return err, nil, nil
  }

  verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
  if err != nil {
    return err, nil, nil
  }

  return nil, verifyKey, signKey
}

type JWTokenServiceSuite struct {
  suite.Suite

  service   *JWTokenService
}

func (s *JWTokenServiceSuite) SetupSuite() {
  verifyKey, signKey, err := initRSAKeys()
  if err != nil {
    panic(err)
  }
  s.service = NewJWTokenService(verifyKey, signKey)
}

func TestJWTokenSuite (t *testing.T) {
  suite.Run(t, &JWTokenServiceSuite{})
}

func (s *JWTokenServiceSuite) TestCreate() {
  tokenString, token, err := s.service.Create(time.Hour)
  if err != nil {
    panic(err)
  }
  assert.NotEqual(s.T(), "", tokenString)
  assert.NotEqual(s.T(), (*jwt.Token)(nil), token)
  //assert.Equal(s.T(), (float64)(3600), token.Claims.(jwt.MapClaims)["exp"].(float64))
}

func (s *JWTokenServiceSuite) TestParse() {
  tokenString, token, err := s.service.Create(time.Hour)
  if err != nil {
    panic(err)
  }

  parsedToken, err := s.service.Parse(tokenString)
  assert.Equal(s.T(), (float64)(token.Claims.(jwt.MapClaims)["exp"].(int64)), parsedToken.Claims.(jwt.MapClaims)["exp"].(float64))
}

func (s *JWTokenServiceSuite) TestValidate() {
  tokenString, _, err := s.service.Create(time.Hour)
  if err != nil {
    panic(err)
  }

  valid := s.service.Validate(tokenString)
  assert.Equal(s.T(), true, valid)
}