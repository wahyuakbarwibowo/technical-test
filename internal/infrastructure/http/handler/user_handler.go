package handler

import (
	"api-golang-codebase/internal/domain/user"

	"github.com/gofiber/fiber/v2"
)

type UserHandler struct {
	userRepo user.Repository
}

func NewUserHandler(userRepo user.Repository) *UserHandler {
	return &UserHandler{userRepo: userRepo}
}

func (h *UserHandler) GetUser(c *fiber.Ctx) error {
	userID := c.Params("id")
	user, err := h.userRepo.GetUserByID(userID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "user not found"})
	}
	return c.JSON(user)
}
