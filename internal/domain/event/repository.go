package event

type Repository interface {
	CreateEvent(event *Event) error
	CreateBooking(event *Booking) error
}
