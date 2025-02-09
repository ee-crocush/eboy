package models

import "game-state/internal/components"

// Unit - сущность юнита.
type Unit struct {
	ID       int64                `json:"id"`
	Health   *components.Health   `json:"health"`
	Mobility *components.Mobility `json:"mobility"`
	Attack   *components.Attack   `json:"attack"`
	//Modifiers []components.Modifier `json:"modifiers"`
}
