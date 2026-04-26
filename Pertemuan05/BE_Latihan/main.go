package main

import (
	router "BE_LATIHAN/Router"
	"BE_LATIHAN/config"
	"BE_LATIHAN/model"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func main() {
	app := fiber.New()
	app.Use(logger.New())
	config.InitDB()

	config.GetDB().AutoMigrate(&model.Mahasiswa{})
	router.SetupRoutes(app)
	app.Listen(":3001")
}
