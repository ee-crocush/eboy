package systems

import (
	"game-state/internal/components"
	"game-state/internal/ecs"
)

// HealthSystem - система для обработки здоровья.
type HealthSystem struct{}

// Update - обновление состояния здоровья сущностей.
func (s *HealthSystem) Update(entities []*ecs.Entity) {
	for _, entity := range entities {
		// Получаем компонент здоровья
		if health, ok := entity.GetComponent("health").(*components.Health); ok {
			// Пример логики: если у сущности здоровье 0, помечаем её как мертвую
			if health.HPNow == 0 {
				// Логика для мертвых сущностей (например, удаление из игры или другие действия)
				// Здесь просто выводим сообщение для примера
				println("Entity", entity.ID, "is dead.")
			}
		}
	}
}
