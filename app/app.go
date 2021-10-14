package app

import (
	"tubespbbo/db/seeds"

	"github.com/gofiber/fiber/v2"
	_ "github.com/joho/godotenv/autoload"
)

var (
	router *fiber.App
)

func seed() {
	var s seeds.Seed
	s.SeedCountry()
}

func StartApplication() {
	seed()
	router = fiber.New()
	mapURLs()
	_ = router.Listen(":8082")
}
