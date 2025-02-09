package systems

import (
	"game-state/internal/ecs"
	"game-state/internal/ecs/components"
)

// HealthSystem - система для обработки здоровья.
type HealthSystem struct{}

// Update - обновление состояния здоровья сущностей.
func (s *HealthSystem) Update(entities []*ecs.Entity) {
	for _, entity := range entities {
		if health, ok := entity.GetComponent("health").(*components.Health); ok {
			if health.HPNow == 0 {
				// Логика для мертвых сущностей
				println("Entity", entity.ID, "is dead.")
			}
		}
	}
}
