package token

import (
	"testing"
	"github.com/stretchr/testify/suite"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/dgrijalva/jwt-go"
)

const testSecret = "testSecret"

type JWTokenServiceHMACSuite struct {
	suite.Suite

	secret []byte

	service   *JWTTokenServiceHMAC
}

func (s *JWTokenServiceHMACSuite) SetupSuite() {
	s.secret = []byte(testSecret)
	s.service = NewJWTTokenServiceHMAC(s.secret)
}

func TestJWTokenServiceHMACSuite (t *testing.T) {
	suite.Run(t, &JWTokenServiceHMACSuite{})
}

func (s *JWTokenServiceHMACSuite) TestCreate() {
	tokenString, err := s.service.Create(time.Hour)
	if err != nil {
		panic(err)
	}
	assert.NotEqual(s.T(), "", tokenString)
	//assert.NotEqual(s.T(), (*jwt.Token)(nil), token)
}

func (s *JWTokenServiceHMACSuite) TestSetClaimValue() {
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

func (s *JWTokenServiceHMACSuite) TestGetClaimValue() {
	tokenString, err := s.service.Create(time.Hour)
	if err != nil {
		panic(err)
	}

	token, err := s.service.Parse(tokenString)
	token.Claims.(jwt.MapClaims)["teststring"] = "test"
	tokenString, err = token.SignedString(s.secret)
	assert.Equal(s.T(), nil, err)

	val, err := s.service.GetClaimValue(tokenString, "teststring")
	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), "test", (val).(string))

	token, err = s.service.Parse(tokenString)
	token.Claims.(jwt.MapClaims)["testint"] = 20
	tokenString, err = token.SignedString(s.secret)
	assert.Equal(s.T(), nil, err)

	val, err = s.service.GetClaimValue(tokenString, "testint")
	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), float64(20), (val).(float64))

	token, err = s.service.Parse(tokenString)
	token.Claims.(jwt.MapClaims)["testfloat"] = 42.5
	tokenString, err = token.SignedString(s.secret)
	assert.Equal(s.T(), nil, err)

	val, err = s.service.GetClaimValue(tokenString, "testfloat")
	assert.Equal(s.T(), nil, err)
	assert.Equal(s.T(), 42.5, (val).(float64))
}

func (s *JWTokenServiceHMACSuite) TestParse() {
	tokenString, err := s.service.Create(time.Hour)
	if err != nil {
		panic(err)
	}

	parsedToken, err := s.service.Parse(tokenString)
	assert.Equal(s.T(), true, parsedToken.Valid)
	//assert.Equal(s.T(), (float64)(token.Claims.(jwt.MapClaims)["exp"].(int64)), parsedToken.Claims.(jwt.MapClaims)["exp"].(float64))
}

func (s *JWTokenServiceHMACSuite) TestValidate() {
	tokenString, err := s.service.Create(time.Hour)
	if err != nil {
		panic(err)
	}

	valid := s.service.Validate(tokenString)
	assert.Equal(s.T(), true, valid)
}