package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var testingPasswords = []string{
	"Heslo123",
	"NiKDoMeE19NeProl29omI",
	"SeJfMeiNName",
	"YoLo",
	"ExpectedAge",
}

const wrongPassword = "JaJaDaDaYesYes"

type BcryptSuite struct {
	suite.Suite
}

func (s *BcryptSuite) TestNewBcrypt() {
	service := NewBcrypt()
	assert.NotEqual(s.T(), nil, service)
}

func (s *BcryptSuite) TestEncryptDecryptCorrect() {
	service := NewBcrypt()

	for _, testPassword := range testingPasswords {
		hash := service.Encrypt(testPassword)
		assert.Equal(s.T(), true, service.Check(hash, testPassword))
	}

}

func (s *BcryptSuite) TestEncryptDecryptIncorrect() {
	service := NewBcrypt()

	for _, testPassword := range testingPasswords {
		hash := service.Encrypt(testPassword)
		assert.NotEqual(s.T(), true, service.Check(hash, wrongPassword))
	}

}

func TestBcryptSuite(t *testing.T) {
	suite.Run(t, &BcryptSuite{})
}
