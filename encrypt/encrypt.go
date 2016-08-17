package encrypt

//BcryptService is an interface for Bcrypt encryption and validation
type BcryptService interface {
	Encrypt(password string) string
	Validate(password string) bool
}
