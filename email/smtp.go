package email

import (
  "net/smtp"
)

// SMTPConfig holds information about smtp server
// that should be used for sending emails.
type SMTPConfig struct {
  Username    string
  Password    string
  ServerHost  string
  ServerPort  string
  SenderAddr  string
}

// SMTPSender is an implementation of emailSender.Service
// interface that uses SMTP
type SMTPSender struct {
  conf  SMTPConfig
  send func(string, smtp.Auth , string, []string, []byte) error
}

// NewSMTPSender is a factory function for SMTPSender
// it uses smtp.SendMail as default sending function
func NewSMTPSender(configuration SMTPConfig) *SMTPSender {
  return &SMTPSender{
    conf:configuration,
    send:smtp.SendMail,
  }
}

// Send will send an email with body given by parameter to
// to addresses given by parameter
func (s *SMTPSender) Send (to []string, body []byte) error {
  addr := s.conf.ServerHost + ":" + s.conf.ServerPort
  auth := smtp.PlainAuth("", s.conf.Username, s.conf.Password, s.conf.ServerHost)
  return s.send(addr, auth, s.conf.SenderAddr, to, body)
}