package email

import "errors"

// Record is a record of an email sent by mock sender
type Record struct {
	To   []string
	Body []byte
}

// MockSender mocks sender implementation and records all Emails
// sent by it
type MockSender struct {
	Used   bool
	Emails []Record
}

// NewMockSender is a factory method for MockSender
func NewMockSender() *MockSender {
	return &MockSender{
		Used:   false,
		Emails: []Record{},
	}
}

// Send will mock sending of an email by recording it in the
// MockSender structure
func (m *MockSender) Send(to []string, body []byte) {
	m.Used = true
	rec := Record{To: to, Body: body}
	m.Emails = append(m.Emails, rec)
}

// GetNumSentEmails will return number of sent emails
func (m *MockSender) GetNumSentEmails() int {
	return len(m.Emails)
}

// GetEmailOnIndex will return nth sent email, n is sent by parameter
// it may return error if index is out of range
func (m *MockSender) GetEmailOnIndex(n int) (Record, error) {
	if n < m.GetNumSentEmails() {
		return Record{}, errors.New("Index out of range")
	}

	return m.Emails[n], nil
}
