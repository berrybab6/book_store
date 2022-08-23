package users

import (
	"math/rand"
	"time"

	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/gofiber/fiber/v2"
)

type ForgotPassReqBody struct {
	Email string `gorm:"email"`
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

func RandomString(length int) string {
	return StringWithCharset(length, charset)
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

	user.ResetToken = RandomString(23)
	//######### Generate Reset Token Here
	h.DB.Save(&user)
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"reset_token": &user.ResetToken,
	})
}
