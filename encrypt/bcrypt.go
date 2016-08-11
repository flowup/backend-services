package encrypt

import "golang.org/x/crypto/bcrypt"

//Bcrypt is a struct
type Bcrypt struct {
}

//NewBcrypt is a constructor creating new Bcrypt
func NewBcrypt() *Bcrypt {
	return &Bcrypt{}
}

// Encrypt is encrypting passwords with bcrypt
func (s *Bcrypt) Encrypt(password string) string {

	// Password hashing with defaultCost of hashing set to be 10
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}
	/*
		// Comparing the password with the hash
		err = bcrypt.CompareHashAndPassword(hashedPassword, []byte(password))
		fmt.Println(err) // nil means it is a match
	*/
	return string(hashedPassword)
}

// Check is comparing non-crypted passwords with encrypted ones
// It returns true is plaintext password is similiar to hashed password
func (s *Bcrypt) Check(hashedPassword, password string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	if err != nil {
		return false
	}
	return true

}
