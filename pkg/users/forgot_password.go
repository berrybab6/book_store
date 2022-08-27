package users

import (
	"math/rand"
	"strings"
	"time"

	"github.com/berrybab6/MovieGo/pkg/common/models"
	"github.com/berrybab6/MovieGo/pkg/common/utils"

	"github.com/gofiber/fiber/v2"
)

// ForgotPassReqBody MovieGo
type ForgotPassReqBody struct {
	Email string `gorm:"email"`
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// StringWithCharset MovieGo
func StringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(b)
}

// RandomString MovieGo
func RandomString(length int) string {
	return StringWithCharset(length, charset)
}

// ForgotPassword godoc
// @Summary      Forgot Password getway
// @Description  change user password
// @Tags         users
// @Accept       json
// @Produce      json
// @Param      	user    body      ForgotPassReqBody  true  "Forgot Password"
// @Success      200      {object}  models.User
// @Failure      400      {object}  httputil.HTTPError
// @Failure      404      {object}  httputil.HTTPError
// @Failure      500      {object}  httputil.HTTPError
// @Router       /users/forgot [put]
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
	var firstName = user.Name

	if strings.Contains(firstName, " ") {
		firstName = strings.Split(firstName, " ")[1]
	}
	emailData := utils.EmailData{
		// URL:       config.Origin + "/forgotPassword/" + user.ResetToken,
		URL:       "http://localhost:8080/users/forgotPassword/" + user.ResetToken,
		FirstName: firstName,
		Subject:   "Your password reset token (valid for 10min)",
	}

	err := utils.SendEmail(&user, &emailData, "resetPassword.html")

	if err != nil {
		return fiber.NewError(fiber.StatusBadGateway, err.Error())

	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"reset_token": &user.ResetToken,
	})
}
