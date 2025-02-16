package modifires

import (
	"game-state/internal/ecs/components/production"
	"time"
)

// ProductivityModifier - модификатор производительности
type ProductivityModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора производительности
func (p *ProductivityModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*production.Production); ok {
		props.Duration = time.Duration(int(float32(props.Duration) * p.Multiplier))
		return props
	}
	return properties
}
