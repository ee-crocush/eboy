package modifires

import (
	"game-state/internal/ecs/components/base"
)

// ArmorModifier - модификатор брони
type ArmorModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора брони
func (a *ArmorModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*base.Health); ok {
		props.Armor = int(float32(props.Armor) * a.Multiplier)
		return props
	}
	return properties
}
