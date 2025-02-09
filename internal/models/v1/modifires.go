package v1

// Modifier - модификатор параметров объектов
type Modifier interface {
	ApplyTo(properties any) any // Применить модификацию
}

// HealthModifier - модификатор здоровья
type HealthModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора здоровья
func (h *HealthModifier) ApplyTo(properties any) any {
	if props, ok := properties.(HealthProperties); ok {
		props.HPNow = int(float32(props.HPNow) * h.Multiplier)
		if props.HPNow > props.HP {
			props.HPNow = props.HP // Ограничение на максимум здоровья
		}
		return props
	}
	return properties
}

// ArmorModifier - модификатор брони
type ArmorModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора брони
func (a *ArmorModifier) ApplyTo(properties any) any {
	if props, ok := properties.(HealthProperties); ok {
		props.Armor = int(float32(props.Armor) * a.Multiplier)
		return props
	}
	return properties
}

// ProductivityModifier - модификатор производительности
type ProductivityModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора производительности
func (p *ProductivityModifier) ApplyTo(properties any) any {
	if props, ok := properties.(ProductionProperties); ok {
		props.Productivity = int(float32(props.Productivity) * p.Multiplier)
		return props
	}
	return properties
}

// SpeedModifier - модификатор скорости
type SpeedModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора скорости
func (s *SpeedModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*MobilityProperties); ok {
		props.Speed = props.Speed * s.Multiplier
		return props
	}
	return properties
}

// VisionModifier - модификатор обзора
type VisionModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора обзора
func (v *VisionModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*MobilityProperties); ok {
		props.Speed = props.Speed * v.Multiplier
		return props
	}
	return properties
}

// DamageModifier - модификатор урона
type DamageModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора урона
func (d *DamageModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*MobilityProperties); ok {
		props.Speed = props.Speed * d.Multiplier
		return props
	}
	return properties
}

// RangeModifier - модификатор дальности атаки
type RangeModifier struct {
	Multiplier float32 `json:"multiplier"` // коэффициент изменения
}

// ApplyTo - применение модификатора дальности атаки
func (r *RangeModifier) ApplyTo(properties any) any {
	if props, ok := properties.(*MobilityProperties); ok {
		props.Speed = props.Speed * r.Multiplier
		return props
	}
	return properties
}
