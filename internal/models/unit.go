package models

import (
	components2 "game-state/internal/ecs/components"
)

// Unit - сущность юнита.
type Unit struct {
	ID       int64                 `json:"id"`
	Health   *components2.Health   `json:"health"`
	Mobility *components2.Mobility `json:"mobility"`
	Attack   *components2.Attack   `json:"attack"`
	//Modifiers []components.Modifier `json:"modifiers"`
}
