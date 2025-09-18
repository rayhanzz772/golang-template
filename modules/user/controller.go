package user

import (
	"belajar-coding/go/modules/user/dto"
	"belajar-coding/go/utils"
	"context"
	"time"

	"github.com/gofiber/fiber/v2"
)

func list(c *fiber.Ctx) error {
	ctxWithTimeout, cancel := context.WithTimeout(c.Context(), 5*time.Second)
	defer cancel()

	perPage, page := utils.GetPaginationParams(map[string]string{
		"per_page": c.Query("per_page"),
		"page":     c.Query("page"),
	}, map[string]int{"per_page": 10, "page": 1})

	users, meta, err := GetUsers(ctxWithTimeout, page, perPage)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Fail(err.Error(), nil))
	}

	return c.JSON(utils.OkPaginate(users, meta, "success"))
}

func create(c *fiber.Ctx) error {
	var req dto.CreateUserRequest

	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Fail(err.Error(), nil))
	}

	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Fail("Validated Error", err.Error()))
	}

	user, err := CreateUser(req)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Fail(err.Error(), nil))
	}

	return c.Status(fiber.StatusCreated).JSON(utils.Ok(user, "User created", nil))
}

func show(c *fiber.Ctx) error {
	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Fail("Invalid UUID", nil))
	}
	user, err := GetUserByID(id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Fail("User not found", nil))
	}
	return c.JSON(user)
}

func update(c *fiber.Ctx) error {

	var req dto.UpdateUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Fail(err.Error(), nil))
	}

	if err := utils.Validate.Struct(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Fail("Validated Error", err.Error()))
	}

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Fail("Invalid UUID", nil))
	}

	if err := utils.CheckUserExist(id); err != nil {
		return c.Status(fiber.StatusNotFound).JSON(utils.Fail("User not found", nil))
	}

	if _, err := UpdateUser(id, req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Fail(err.Error(), nil))
	}

	return c.JSON(utils.SuccessResponseNoData("User updated"))
}

func delete(c *fiber.Ctx) error {

	id, err := utils.GetUUIDParam(c, "id")
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(utils.Fail("Invalid UUID", nil))
	}
	if err := DeleteUser(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(utils.Fail(err.Error(), nil))
	}
	return c.JSON(utils.Ok("User deleted", "success", nil))
}
