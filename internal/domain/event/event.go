package event

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	gorm.Model
	Description string
	Quota       uint
	StartDate   time.Time
	EndDate     time.Time
}

type Booking struct {
	gorm.Model
	Description string
	Quota       uint
	StartDate   time.Time
	EndDate     time.Time
}
