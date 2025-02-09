package systems

import (
	"game-state/internal/components"
	"game-state/internal/ecs"
)

// CombatSystem - система для обработки боевых действий.
type CombatSystem struct{}

// Update - обновление состояния боевых действий.
func (s *CombatSystem) Update(entities []*ecs.Entity) {
	for _, entity := range entities {
		if attack, ok := entity.GetComponent("attack").(*components.Attack); ok {
			// Пример логики: если это дальнобойная атака, проверяем возможную цель
			if attack.IsRange {
				// Логика дальнобойных атак
			} else {
				// Логика ближнего боя
			}
		}
	}
}
