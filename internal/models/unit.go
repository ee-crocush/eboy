package models

import (
	"game-state/internal/ecs/components/base"
	"game-state/internal/ecs/components/combat"
	components "game-state/internal/ecs/components/mobility"
)

// Unit - сущность юнита.
type Unit struct {
	ID       int64                `json:"id"`       // ID юнита
	Name     string               `json:"name"`     // Название юнита
	Health   *base.Health         `json:"health"`   // Здоровье
	Mobility *components.Mobility `json:"mobility"` // Перемещение
	Attack   *combat.Attack       `json:"attack"`   // Атака
	//Modifiers []components.Modifier `json:"modifiers"`
}
