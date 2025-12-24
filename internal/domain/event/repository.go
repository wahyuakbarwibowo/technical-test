package event

type Repository interface {
	CreateEvent(event *Event) error
}
