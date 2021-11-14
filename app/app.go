package app

import (
	"tubespbbo/db/seeds"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	_ "github.com/joho/godotenv/autoload"
)

var (
	router *fiber.App
)

func seed() {
	var s seeds.Seed
	s.SeedPM()
	s.SeedUser()
	s.SeedProduct()
	s.SeedTransaction()
	s.SeedPayment()
	s.SeedTD()
	s.SeedExpence()
}

func StartApplication() {
	seed()
	router = fiber.New()
	router.Use(cors.New())
	mapURLs()
	_ = router.Listen(":8082")
}
