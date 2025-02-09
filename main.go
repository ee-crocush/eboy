package main

import (
	"fmt"
	handlers "game-state/internal/api/http"
	"game-state/internal/api/websocket"
	"game-state/internal/ecs"
	"game-state/internal/ecs/components"
	"game-state/internal/ecs/systems"
	"github.com/gofiber/fiber/v3"
	"log"
)

func main() {
	initWorld()
	app := fiber.New()

	// Обработчик для POST-запроса на /get_state
	app.Post("/get_init_state", handlers.GetInitWorldState)

	if err := app.Listen(":3000"); err != nil {
		log.Fatalf("Error starting Fiber server: %v", err)
	}

	// Запуск WebSocket сервера на порту 8080
	if err := websocket.Start(":8080"); err != nil {
		log.Fatalf("Error starting WebSocket server: %v", err)
	}
}

func initWorld() {
	// Создаем мир
	world := ecs.NewWorld()

	// Создаем сущность юнита
	unit := ecs.NewEntity(1)

	// Добавляем компоненты юнита
	unit.AddComponent("health", &components.Health{HP: 100, HPNow: 100, Armor: 10})
	unit.AddComponent("mobility", &components.Mobility{Speed: 5, Vision: 3})
	unit.AddComponent("attack", &components.Attack{Damage: 20, IsRange: false, Range: 1})

	// Добавляем сущность в мир
	world.AddEntity(unit)

	// Создаем и добавляем системы
	healthSystem := &systems.HealthSystem{}
	movementSystem := &systems.MovementSystem{}
	combatSystem := &systems.CombatSystem{}

	world.AddSystem(healthSystem)
	world.AddSystem(movementSystem)
	world.AddSystem(combatSystem)

	// Обновляем мир (и все системы) каждый цикл игры
	world.Update()

	// Выводим состояние юнита
	fmt.Println("Unit Health:", unit.GetComponent("health"))
	fmt.Println("Unit Mobility:", unit.GetComponent("mobility"))
	fmt.Println("Unit Attack:", unit.GetComponent("attack"))
}
