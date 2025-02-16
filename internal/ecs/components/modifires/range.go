package modifires

import "game-state/internal/ecs/components/combat"

// RangeModifier - модификатор дальности атаки
type RangeModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора дальности атаки
func (r *RangeModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*combat.Attack); ok {
		props.Range = byte(float32(props.Range) * r.Multiplier)
		return props
	}
	return properties
}
