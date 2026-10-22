package main

import (
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v3"
)

type Response struct {
	Message   string `json:"message"`
	Timestamp int64  `json:"timestamp"`
}

func main() {
	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.JSON(Response{
			Message:   "My name is Tariq Kadir",
			Timestamp: time.Now().UnixMilli(),
		})
	})

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}
	log.Fatal(app.Listen(":" + port))
}
