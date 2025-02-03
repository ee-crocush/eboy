package game_state

import (
	handlers "game-state/internal/api/http"
	"game-state/internal/api/websocket"
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	app := fiber.New()

	// Обработчик для POST-запроса на /get_state
	app.Post("/get_init_state", handlers.GetInitWorldState)

	app.Listen(":3000")

	// Запуск WebSocket сервера на порту 8080
	if err := websocket.Start(":8080"); err != nil {
		log.Fatalf("Error starting WebSocket server: %v", err)
	}
}
