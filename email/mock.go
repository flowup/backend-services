package email

// Record is a record of an email sent by mock sender
type Record struct {
	To   []string
	Body []byte
}

// MockSender mocks sender implementation and records all Emails
// sent
type MockSender struct {
<<<<<<< HEAD
	Emails chan Record
=======
	Emails            chan Record
	maxRecordedEmails int

>>>>>>> 0058bdb83826f991fc52e591708bebbfe662ae1f
}

// NewMockSender is a factory method for MockSender
func NewMockSender() *MockSender {
	return &MockSender{
		Emails: make(chan Record),
	}
}

// NewMockSenderWithMax is a factory method for MockSender that
// also sets maximum of stored emails to value given by parameter
func NewMockSenderWithMax(MaxEmails int) *MockSender {
	return &MockSender{
		Emails: make(chan Record, MaxEmails),
	}
}

// Send will mock sending of an email by recording it in the
// MockSender structure
func (m *MockSender) Send(to []string, body []byte) error {
	m.Emails <- Record{To: to, Body: body}

	return nil
}

// GetNumSentEmails will return number of sent emails
func (m *MockSender) GetNumSentEmails() int {
	return len(m.Emails)
}

// GetEmail will return email on top of the recording channel
// this will block the program if no email is present
func (m *MockSender) GetEmail() Record {
	return <-m.Emails
}

// SetMaxEmailsAndReset will set a maximum of stored emails
// to value given by parameter and reset stored emails
func (m *MockSender) SetMaxEmailsAndReset(MaxEmails int) {
	close(m.Emails)
	m.Emails = make(chan Record, MaxEmails)
}

// Reset will delete all emails present in MockSender
func (m *MockSender) Reset() {
	close(m.Emails)
	m.Emails = make(chan Record, cap(m.Emails))
}
