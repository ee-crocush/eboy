package modifires

import "game-state/internal/ecs/components/combat"

// DamageModifier - модификатор урона
type DamageModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора урона
func (d *DamageModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*combat.Attack); ok {
		props.Damage = int(float32(props.Damage) * d.Multiplier)
		return props
	}
	return properties
}
