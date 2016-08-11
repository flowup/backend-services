package encrypt

//BcryptService of and passawordService
type BcryptService interface {
	Encrypt(password string) string
	Validate(password string) bool
}
