package encrypt

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

// testingPasswords are passwords that will be hashed by encryption function
var testingPasswords = []string{
	"Heslo123",
	"NiKDoMeE19NeProl29omI",
	"SeJfMeiNName",
	"YoLo",
	"ExpectedAge",
}

const wrongPassword = "JaJaDaDaYesYes"

// BcryptSuite defines suite
type BcryptSuite struct {
	suite.Suite
}

// TestNewBcrypt tests if Bcrypt service is created correctly
func (s *BcryptSuite) TestNewBcrypt() {
	service := NewBcrypt()
	assert.NotEqual(s.T(), nil, service)
}

/* TestEncryptDecryptCorrect is testing correct hashing and validating. It creates hashes from plain text passwords
and than validating these hashes by compering them with their plain text versions*/
func (s *BcryptSuite) TestEncryptDecryptCorrect() {
	service := NewBcrypt()

	for _, testPassword := range testingPasswords {
		hash := service.Encrypt(testPassword)
		assert.Equal(s.T(), true, service.Check(hash, testPassword))
	}

}

// TestEncryptDecryptIncorrect is hashing passwords and then trying to compare these hashes with different password
func (s *BcryptSuite) TestEncryptDecryptIncorrect() {
	service := NewBcrypt()

	for _, testPassword := range testingPasswords {
		hash := service.Encrypt(testPassword)
		assert.NotEqual(s.T(), true, service.Check(hash, wrongPassword))
	}

}

//	Passing of our suite (BcryptSuite) to suite.Run to be able to run the tests
func TestBcryptSuite(t *testing.T) {
	suite.Run(t, &BcryptSuite{})
}
