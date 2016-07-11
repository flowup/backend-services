package email

import (
  "github.com/stretchr/testify/suite"
  "net/smtp"
  "github.com/stretchr/testify/assert"
  "testing"
)

type SMTPSenderSuite struct {
  suite.Suite

  service *SMTPSender
  recorder *emailRecorder
}

type emailRecorder struct {
  addr string
  auth smtp.Auth
  from string
  to   []string
  msg  []byte
}

func mockSend(errToReturn error) (func(string, smtp.Auth, string, []string, []byte) error, *emailRecorder) {
  r := new(emailRecorder)
  return func(addr string, a smtp.Auth, from string, to []string, msg []byte) error {
    *r = emailRecorder{addr, a, from, to, msg}
    return errToReturn
  }, r
}

func (s *SMTPSenderSuite) SetupSuite() {
  mock, rec := mockSend(nil)
  s.service = &SMTPSender{send : mock}
  s.recorder = rec
}

func (s *SMTPSenderSuite) TestSend() {
  err := s.service.Send([]string{"test@test.com"},[]byte("Hello"))

  assert.Equal(s.T(), nil, err)
  assert.Equal(s.T(), "Hello", string(s.recorder.msg))
}

func TestSMTPSenderSuite (t *testing.T) {
  suite.Run(t, &SMTPSenderSuite{})
}