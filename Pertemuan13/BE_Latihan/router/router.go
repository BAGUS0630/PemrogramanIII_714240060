package router

import (
	"BE_LATIHAN/config/middleware"
	"BE_LATIHAN/handler"
	"BE_LATIHAN/model"

	"github.com/gofiber/fiber/v2"
	swagger "github.com/gofiber/swagger"

	_ "BE_LATIHAN/docs"
)

func SetupRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.JSON(model.Response{
			Message: "API be_latihan aktif",
		})
	})

	app.Get("/swagger/*", swagger.HandlerDefault)
	app.Get("/docs", func(c *fiber.Ctx) error {
		return c.Redirect("/swagger/index.html")
	})

	app.Post("/register", handler.Register)
	app.Post("/login", handler.Login)
	app.Post("/api/change-password", middleware.JWTProtected(""), handler.ChangePassword)

	mahasiswa := app.Group("/api/mahasiswa")
	mahasiswa.Get("/", middleware.JWTProtected(""), handler.GetAllMahasiswa)
	mahasiswa.Get("/:npm", middleware.JWTProtected("admin"), handler.GetMahasiswaByNPM)
	mahasiswa.Post("/", middleware.JWTProtected("admin"), handler.InsertMahasiswa)
	mahasiswa.Put("/:npm", middleware.JWTProtected("admin"), handler.UpdateMahasiswa)
	mahasiswa.Delete("/:npm", middleware.JWTProtected("admin"), handler.DeleteMahasiswa)
}