package main

import (
	"BE_LATIHAN/config"
	"BE_LATIHAN/model"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	config.InitDB()

	config.GetDB().AutoMigrate(&model.Mahasiswa{})

	app.Listen("3000")
}
