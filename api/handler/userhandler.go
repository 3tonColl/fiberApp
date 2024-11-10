package handler

import (
	"context"
	"fmt"
	"github.com/3tonColl/fiberApp/fiberApp/internal/controller"
	"github.com/3tonColl/fiberApp/fiberApp/model"
	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	GetUser(c *fiber.Ctx) error
	InsertUser(c *fiber.Ctx) error
	UpdateUser(c *fiber.Ctx) error
	DeleteUser(c *fiber.Ctx) error
	GetUserPassword(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
}
type userHandler struct {
	ctx        context.Context
	controller controller.UserController
}

func (us *userHandler) InsertUser(c *fiber.Ctx) error {
	var User model.User
	if err := c.BodyParser(&User); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Cannot parse JSON"})
	}
	if err := us.controller.InsertUser(c.Context(), User); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": fmt.Sprintf("%s", err)})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "successfully"})
}
