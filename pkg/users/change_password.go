package users

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type ChangePassReqBody struct {
	OldPassword string `gorm:"oldpassword"`
	Password    string `gorm:"password"`
}

func (h handler) ChangePassword(c *fiber.Ctx) error {
	req := ChangePassReqBody{}
	id := c.Params("id")

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var user models.User

	record := h.DB.Where("id = ?", id).First(&user)
	if record.Error != nil {
		c.Status(fiber.StatusNotFound)
		return c.JSON(fiber.Map{
			"message": "User with id Doesnot exist",
		})
	}

	if err := user.CheckPassword(req.OldPassword); err != nil {
		c.Status(fiber.StatusBadRequest)
		return c.JSON(fiber.Map{
			"message": "Incorrect Password",
		})
	}

	if err := user.HashPassword(req.Password); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	h.DB.Save(&user)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Password changed succesfully",
	})

}
