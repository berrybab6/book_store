package users

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type ForgotPassReqBody struct {
	Email string `gorm:"email"`
}

func (h handler) ForgotPassword(c *fiber.Ctx) error {
	req := ForgotPassReqBody{}

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}
	var user models.User

	if resp := h.DB.Where("email=?", req.Email).First(&user); resp.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, resp.Error.Error())
	}

	//######### Generate Reset Token Here

	return nil
}
