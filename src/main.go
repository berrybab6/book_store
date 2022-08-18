package main

import (
	"log"

	"github.com/berrybab6/MovieGo/pkg/books"
	"github.com/berrybab6/MovieGo/pkg/common/config"
	"github.com/berrybab6/MovieGo/pkg/common/db"
	"github.com/gofiber/fiber/v2"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at Config", err)
	}
	app := fiber.New()
	db := db.Init(c.DBUrl)
	// app.Get("/", func(ctx *fiber.Ctx) error {
	// 	return ctx.Status(fiber.StatusOK).SendString(c.Port)
	// })
	books.RegisterRoutes(app, db)
	app.Listen(c.Port)

}
