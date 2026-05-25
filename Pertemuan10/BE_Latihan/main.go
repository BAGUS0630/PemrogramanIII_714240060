package main

import (
	router "BE_LATIHAN/Router"
	"BE_LATIHAN/config"
	"BE_LATIHAN/model"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	config.InitDB()

	config.GetDB().AutoMigrate(&model.Mahasiswa{})

	// Basic CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: strings.Join(config.GetAllowedOrigins(), ","),
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
	}))
	router.SetupRoutes(app)

	app.Listen(":3000")

}
