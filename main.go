package main

import (
	handlers "game-state/internal/api/http"
	"game-state/internal/api/websocket"
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	app := fiber.New()

	// Инициализация маршрутов
	handlers.SetupRoutes(app)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting Fiber server: %v", err)
	}

	// Запуск WebSocket сервера на порту 8080
	if err := websocket.Start(":8080"); err != nil {
		log.Fatalf("Error starting WebSocket server: %v", err)
	}
}
