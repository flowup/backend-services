## Encrypt

So far Encrypt service uses just Bcrypt that is implementing Provos and Mazières's bcrypt adaptive hashing algorithm. This service allows you to make hash from your plain-text password and than compare this bcrypt hash with its possible plain-text equivalent. This service offers two methods Encrypt and Check as are shown below. 

## Code Example

Following part of code is showing how you can import and use Encrypt service in order to hash your password:

    import "flowdock.eu/flowup/services/encrypt"

    encrypt := encrypt.NewBcrypt()

    password := "MyUltimateSecret"

    hashedPassword := encrypt.Encrypt(password)

 If you than have a password that is hashed via Encrypting service or with Bcrypt itself you are able to compare that password with its possible plain-text form:

    if encrypt.Check(hashedPassword, password){
        // True if matches
    } 

## Tests

All tests can be executed by running: 

    goconvey 

in the folder with test file, alternatively you can use command: 
	
    go test

## License