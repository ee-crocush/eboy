package modifires

import "game-state/internal/ecs/components/mobility"

// SpeedModifier - модификатор скорости
type SpeedModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора скорости
func (s *SpeedModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*mobility.Mobility); ok {
		props.Speed = props.Speed * s.Multiplier
		return props
	}
	return properties
}
