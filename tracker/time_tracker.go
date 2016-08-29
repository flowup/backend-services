package tracker

//Tracker is an interface for Bcrypt encryption and validation
type Tracker interface {
	StatNew() *Event
	Start(e *Event)
	Stop(e *Event)
	Total(e *Event) int64
}
