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
			// Логика атаки
			if attack.IsRange {
				println("Entity", entity.ID, "attack with range damage:", attack.Damage, "range:", attack.Range)
			} else {
				println("Entity", entity.ID, "attack with melee damage:", attack.Damage)
			}
		}
	}
}
