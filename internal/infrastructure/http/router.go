package http

import (
	"api-golang-codebase/internal/infrastructure/http/handler"
	"api-golang-codebase/internal/infrastructure/presistence/postgres"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterRoutes(app *fiber.App, db *gorm.DB) {
	userRepo := postgres.NewUserRepository(db)
	eventRepo := postgres.NewEventRepository(db)

	userHandler := handler.NewUserHandler(userRepo)
	eventHandler := handler.NewEventHandler(eventRepo)

	app.Get("/user/:id", userHandler.GetUser)
	app.Post("/event", eventHandler.CreateEvent)
}
