package systems

import (
	"game-state/internal/components"
	"game-state/internal/ecs"
)

// MovementSystem - система для обработки перемещения.
type MovementSystem struct{}

// Update - обновление состояния перемещения сущностей.
func (s *MovementSystem) Update(entities []*ecs.Entity) {
	for _, entity := range entities {
		// Получаем компонент мобильности
		if mobility, ok := entity.GetComponent("mobility").(*components.Mobility); ok {
			// Логика перемещения сущности, например, если скорость больше 0, перемещаем её
			// Здесь просто выводим пример перемещения
			println("Entity", entity.ID, "moving at speed:", mobility.Speed)
		}
	}
}
