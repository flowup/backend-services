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
  signBytes, err := ioutil.ReadFile("test_fixture/app.rsa")
  if err != nil {
    return nil, nil, err
  }
  signKey, err := jwt.ParseRSAPrivateKeyFromPEM(signBytes)
  if err != nil {
    return nil, nil, err
  }
  verifyBytes, err := ioutil.ReadFile("test_fixture/app.rsa.pub")
  if err != nil {
    return nil, nil, err
  }

  verifyKey, err := jwt.ParseRSAPublicKeyFromPEM(verifyBytes)
  if err != nil {
    return nil, nil, err
  }

  return verifyKey, signKey, nil
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
  tokenString, err := s.service.Create(time.Hour)
  if err != nil {
    panic(err)
  }
  assert.NotEqual(s.T(), "", tokenString)
  //assert.NotEqual(s.T(), (*jwt.Token)(nil), token)
}

func (s *JWTokenServiceSuite) TestSetClaimValue() {
  tokenString, err := s.service.Create(time.Hour)
  if err != nil {
    panic(err)
  }

  tokenString, err = s.service.SetClaimValue(tokenString, "testclaim", "testval")
  token, err := s.service.Parse(tokenString)
  assert.Equal(s.T(), nil, err)
  assert.Equal(s.T(), "testval", token.Claims.(jwt.MapClaims)["testclaim"].(string))

  tokenString, err = s.service.SetClaimValue(tokenString, "testint", 42)
  token, err = s.service.Parse(tokenString)
  assert.Equal(s.T(), nil, err)
  assert.Equal(s.T(), float64(42), token.Claims.(jwt.MapClaims)["testint"].(float64))

  tokenString, err = s.service.SetClaimValue(tokenString, "testfloat", 42.5)
  token, err = s.service.Parse(tokenString)
  assert.Equal(s.T(), nil, err)
  assert.Equal(s.T(), float64(42.5), token.Claims.(jwt.MapClaims)["testfloat"].(float64))
}

func (s *JWTokenServiceSuite) TestGetClaimValue() {
  tokenString, err := s.service.Create(time.Hour)
  if err != nil {
    panic(err)
  }

  token, err := s.service.Parse(tokenString)
  token.Claims.(jwt.MapClaims)["teststring"] = "test"
  tokenString, err = token.SignedString(s.service.SignKey)
  assert.Equal(s.T(), nil, err)

  val, err := s.service.GetClaimValue(tokenString, "teststring")
  assert.Equal(s.T(), nil, err)
  assert.Equal(s.T(), "test", (val).(string))

  token, err = s.service.Parse(tokenString)
  token.Claims.(jwt.MapClaims)["testint"] = 20
  tokenString, err = token.SignedString(s.service.SignKey)
  assert.Equal(s.T(), nil, err)

  val, err = s.service.GetClaimValue(tokenString, "testint")
  assert.Equal(s.T(), nil, err)
  assert.Equal(s.T(), float64(20), (val).(float64))

  token, err = s.service.Parse(tokenString)
  token.Claims.(jwt.MapClaims)["testfloat"] = 42.5
  tokenString, err = token.SignedString(s.service.SignKey)
  assert.Equal(s.T(), nil, err)

  val, err = s.service.GetClaimValue(tokenString, "testfloat")
  assert.Equal(s.T(), nil, err)
  assert.Equal(s.T(), 42.5, (val).(float64))
}

func (s *JWTokenServiceSuite) TestParse() {
  tokenString, err := s.service.Create(time.Hour)
  if err != nil {
    panic(err)
  }

  parsedToken, err := s.service.Parse(tokenString)
  assert.Equal(s.T(), true, parsedToken.Valid)
  //assert.Equal(s.T(), (float64)(token.Claims.(jwt.MapClaims)["exp"].(int64)), parsedToken.Claims.(jwt.MapClaims)["exp"].(float64))
}

func (s *JWTokenServiceSuite) TestValidate() {
  tokenString, err := s.service.Create(time.Hour)
  if err != nil {
    panic(err)
  }

  valid := s.service.Validate(tokenString)
  assert.Equal(s.T(), true, valid)
}