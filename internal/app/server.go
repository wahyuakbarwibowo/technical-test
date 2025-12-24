package app

import (
	"api-golang-codebase/config"
	"api-golang-codebase/internal/infrastructure/http"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Server struct {
	app *fiber.App
	db  *gorm.DB
}

func NewServer(cfg *config.Config, db *gorm.DB) *Server {
	app := fiber.New()
	http.RegisterRoutes(app, db)
	return &Server{app: app, db: db}
}

func (s *Server) Run() error {
	return s.app.Listen("localhost:8080")
}
