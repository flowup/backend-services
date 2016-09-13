package tracker

//Tracker is an interface for time tracking service
type Tracker interface {
	StatNew() *Event
	Start(e *Event)
	Stop(e *Event)
	Total(e *Event) int64
}
