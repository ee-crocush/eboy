package modifires

import "game-state/internal/ecs/components/mobility"

// VisionModifier - модификатор обзора
type VisionModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора обзора
func (v *VisionModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*mobility.Mobility); ok {
		props.Vision = byte(float32(props.Vision) * v.Multiplier)
		return props
	}
	return properties
}
