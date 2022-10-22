package controllers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/harikesh-yadav/gofiber_api/helpers"
	"github.com/harikesh-yadav/gofiber_api/models"
	"gopkg.in/validator.v2"
)

func Registration(ctx *fiber.Ctx) error {
	user := &models.User{}

	if err := ctx.BodyParser(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	user.ID = uuid.New()
	user.Created_at = time.Now()

	if err := validator.Validate(user); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err,
		})
	}

	err := helpers.CreateUser(user)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"error": false,
		"msg":   "User successfully created",
	})
}

func Login(ctx *fiber.Ctx) error {

	credential := &models.Credential{}
	if err := ctx.BodyParser(credential); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	if err := validator.Validate(credential); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	result, err := helpers.Login(credential)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return ctx.JSON(fiber.Map{
		"error":  false,
		"result": result,
	})
}
