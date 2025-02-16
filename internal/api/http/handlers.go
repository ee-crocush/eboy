package http

import (
	"fmt"
	"game-state/internal/ecs"
	"game-state/internal/ecs/components/base"
	"game-state/internal/ecs/components/combat"
	"game-state/internal/ecs/components/mobility"
	"game-state/internal/ecs/systems"
	"game-state/internal/models"
	"github.com/gofiber/fiber/v3"
	"strconv"
)

// GenerateWorldResponse - Ответ на запрос на генерацию мира
type GenerateWorldResponse struct {
	Area struct {
		ID     int `json:"id"`
		Radius int `json:"radius"`
	} `json:"area"`
	Objects []interface{} `json:"objects"`
}

func SetupRoutes(app *fiber.App) {
	// Регистрируем POST-ручку
	app.Post("/generate_world", GenerateWorld)
}

func GenerateWorld(c fiber.Ctx) error {
	user_id := c.FormValue("user_id")

	fmt.Printf("Generating world for user %d\n", user_id)

	// Создаем мир
	world := ecs.NewWorld()

	units := []ecs.Entity{}

	for i := 0; i < 5; i++ {
		unit := ecs.NewEntity(int64(i), "Unit"+strconv.Itoa(i))
		unit.AddComponent("health", &base.Health{HP: 100, HPNow: 100, Armor: 10})
		unit.AddComponent("mobility", &mobility.Mobility{Speed: 5, Vision: 3})
		unit.AddComponent("attack", &combat.Attack{Damage: 20, IsRange: false, Range: 1})

		world.AddEntity(unit)
		units = append(units, *unit)
	}

	// Создаем и добавляем системы
	healthSystem := &systems.HealthSystem{}
	movementSystem := &systems.MovementSystem{}
	combatSystem := &systems.CombatSystem{}

	world.AddSystem(healthSystem)
	world.AddSystem(movementSystem)
	world.AddSystem(combatSystem)

	// Обновляем мир (и все системы) каждый цикл игры
	world.Update()

	// Формируем список объектов (сущностей)
	var objects []interface{}

	for _, unit := range units {
		unitData := &models.Unit{
			ID:       unit.ID,
			Health:   unit.GetComponent("health").(*base.Health),
			Mobility: unit.GetComponent("mobility").(*mobility.Mobility),
			Attack:   unit.GetComponent("attack").(*combat.Attack),
		}
		objects = append(objects, unitData)
	}

	// Формируем ответ
	return c.JSON(GenerateWorldResponse{
		Area: struct {
			ID     int `json:"id"`
			Radius int `json:"radius"`
		}{ID: 1, Radius: 50},
		Objects: objects,
	})
}
