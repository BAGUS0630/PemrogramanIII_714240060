package router

import (
	"BE_LATIHAN/config/middleware"
	"BE_LATIHAN/handler"
	"BE_LATIHAN/model"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(model.Response{
			Message: "API be_latihan aktif",
		})
	})

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)

	mahasiswa := app.Group("/api/mahasiswa", middleware.JWTProtected("admin"))
	mahasiswa.Get("/", handler.GetAllMahasiswa)
	mahasiswa.Get("/:npm", handler.GetMahasiswaByNPM)
	mahasiswa.Post("/", handler.InsertMahasiswa)
	mahasiswa.Put("/:npm", handler.UpdateMahasiswa)
	mahasiswa.Delete("/:npm", handler.DeleteMahasiswa)
}