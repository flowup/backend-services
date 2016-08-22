package timeTracker

//TimeTrackingService is an interface for Bcrypt encryption and validation
type TimeTrackingService interface {
	StatNew() *Event
	Start(e *Event)
	Stop(e *Event)
}
