package postgres

import (
	"api-golang-codebase/internal/domain/event"

	"gorm.io/gorm"
)

type EventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) *EventRepository {
	return &EventRepository{db: db}
}

func (r *EventRepository) CreateEvent(event *event.Event) error {
	return r.db.Create(event).Error
}
