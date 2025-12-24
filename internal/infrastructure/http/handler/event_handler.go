package handler

import (
	"api-golang-codebase/internal/domain/event"

	"github.com/gofiber/fiber/v2"
)

type EventHandler struct {
	eventRepo event.Repository
}

func NewEventHandler(eventRepo event.Repository) *EventHandler {
	return &EventHandler{eventRepo: eventRepo}
}

func (h *EventHandler) CreateEvent(c *fiber.Ctx) error {
	e := new(event.Event)

	if err := c.BodyParser(e); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": "bad request"})
	}
	err := h.eventRepo.CreateEvent(e)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "success to create event"})
}
