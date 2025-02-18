package systems

import (
	"game-state/internal/ecs"
	"game-state/internal/ecs/components/mobility"
)

// MovementSystem - система для обработки перемещения.
type MovementSystem struct{}

// Update - обновление состояния перемещения сущностей.
func (s *MovementSystem) Update(entities []*ecs.Entity) {
	for _, entity := range entities {
		if mobility, ok := entity.GetComponent("mobility").(*mobility.Mobility); ok {
			// Логика перемещения
			println("Entity", entity.ID, "moving at speed:", mobility.Speed)
		}
	}
}
