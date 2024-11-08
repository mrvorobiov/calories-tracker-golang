package main

import (
	"app/internal/db"
	"app/internal/meal"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	app := fiber.New()

	db, err := db.Open()

	if err != nil {
		log.Fatal(err)
	}

	meal.InitRouter(db, app)

	app.Listen(":80")
}
