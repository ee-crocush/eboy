package modifires

import (
	"game-state/internal/ecs/components/base"
)

// HealthModifier - модификатор здоровья
type HealthModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора здоровья
func (h *HealthModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*base.Health); ok {
		props.HPNow = int(float32(props.HPNow) * h.Multiplier)
		if props.HPNow > props.HP {
			props.HPNow = props.HP
		}
		return props
	}
	return properties
}
