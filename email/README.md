## Email

Email service is service allowing user to easily send emails. For that purpose this service contains only one method `Send`.

## Code Example

First config file needs to be set that means that you have to fill email.SMTPConfing and this config is than used to initialize email service:
    
    import "flowdock.eu/flowup/services/email"

    config := email.SMTPConfig{
		Username:   "example@example.com",
		Password:   "example123",
		ServerHost: "smtp.example.com",
		ServerPort: "587",                        //gmail ServerPort
		SenderAddr: "example@example.com",
	}
    
    email := email.NewSMTPSender(config)

Email `SENDING` is than very simple: 
    
    err = c.email.Send([]string{"exampleDestination@example.com"}, []byte("Your email message")) 

## Tests

All tests can be executed by running: 

    goconvey 

in the folder with test file, alternatively you can use command: 

    go test

## License