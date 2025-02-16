package modifires

import (
	"game-state/internal/ecs/components/production"
)

// GatheringModifier - модификатор добычи
type GatheringModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора добычи
func (p *GatheringModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*production.Gathering); ok {
		props.Productivity = int(float32(props.Productivity) * p.Multiplier)
		return props
	}
	return properties
}
