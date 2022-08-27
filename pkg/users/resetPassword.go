package users

import (
	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

// ResetPasswordInput MovieGo
type ResetPasswordInput struct {
	Password        string `json:"password" binding:"required"`
	PasswordConfirm string `json:"passwordConfirm" binding:"required"`
	Email           string `json:"email" binding:"required"`
}

// ResetPassword godoc
// @Summary      Reset Password getway
// @Description  reset user password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param        resettoken   path      string  true  "User ResetToken"
// @Param      	user    body      ResetPasswordInput  true  "Reset Password"
// @Success      200      {object}  models.User
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /users/reset/{resettoken} [put]
func (h handler) ResetPassword(c *fiber.Ctx) error {

	req := ResetPasswordInput{}
	resetToken := c.Params("resettoken")

	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())
	}

	if req.Password != req.PasswordConfirm {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Passwords do not match"})

		return err
	}
	var user *models.User

	if resp := h.DB.Where("email=?", req.Email).First(&user); resp.Error != nil {
		return fiber.NewError(fiber.StatusBadRequest, resp.Error.Error())
	}

	if resetToken == "" {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Reset Token Empty"})

		return err
	}
	if user.ResetToken != resetToken || resetToken == "" {
		err := c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"status": "fail", "message": "Reset Token do not match"})

		return err
	}
	if err := user.HashPassword(req.Password); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, err.Error())

	}

	user.ResetToken = ""
	h.DB.Save(&user)
	// hashedPassword, _ := HashPassword(userCredential.Password)

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "user Password changed!!!",
	})
}
