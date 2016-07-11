package email

// Service is an interface for implementations
// of email sending services
type Sender interface {
  Send(to []string, body []byte)
}