package auth

import (
	"belajar-coding/go/modules/auth/dto"
	"belajar-coding/go/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

func LoginHandler(c *fiber.Ctx) error {
	var req dto.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(utils.Fail(err.Error(), nil))
	}

	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(400).JSON(utils.Fail(err.Error(), nil))
	}

	token, err := ServiceLogin(req)
	if err != nil {
		return c.Status(401).JSON(utils.Fail("invalid email or passwords", nil))
	}

	return c.JSON(utils.Ok(fiber.Map{
		"token": token,
	}, "Login success", nil))
}

func MeHandler(c *fiber.Ctx) error {
	userID, ok := c.Locals("userID").(string)
	if !ok || userID == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Fail("Unauthorized", nil))
	}

	id, err := uuid.Parse(userID)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(utils.Fail("Invalid token subject", nil))
	}

	user, err := FindUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Fail("User not found", nil))
	}

	return c.JSON(utils.Ok(user, "success", nil))
}
