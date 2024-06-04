package main

import (
	"muhammadiyah/Config"
	"muhammadiyah/Route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	app := fiber.New()

	if err := godotenv.Load(".env"); err != nil {
		panic(err)
	}

	app.Use(cors.New(cors.Config{
		AllowHeaders: "Origin,Content-Type,Accept,User-Agent,Content-Length,Authorization",
		AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,PUT,DELETE,PATCH,OPTIONS",
	}))
	// Inisialisasi database
	Config.InitDB()

	// Menggunakan router sebagai handler pada root route "/api"
	Route.Router(app)
	// Menjalankan server pada port 3000
	err := app.Listen(":3025")
	if err != nil {
		panic(err)
	}

}
