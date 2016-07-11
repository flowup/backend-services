package emailSender

// Service is an interface for implementations
// of email sending services
type Service interface {
  Send(to []string, body []byte)
}