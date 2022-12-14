package users

import (
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// handler MovieGo
type handler struct {
	DB *gorm.DB
}

// RegisterUserRoutes godoc
func RegisterUserRoutes(app *fiber.App, db *gorm.DB) {
	h := &handler{
		DB: db,
	}

	routes := app.Group("users")
	routes.Post("/", h.CreateUser)
	routes.Post("/login", h.LoginUser)
	routes.Put("/forgot", h.ForgotPassword)
	routes.Put("/reset/:resettoken", h.ResetPassword)

	secured := app.Group("/secured").Use(ValidateJWT)
	{
		secured.Get("/", h.GetUsers)
		// secured.Put("/:id", h.ChangePassword)

		// secured.GET("/ping", controllers.Ping)
	}
}
