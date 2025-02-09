package components

import (
	"game-state/internal/models/v2/errors"
)

// HealthType определяет тип сущности с здоровьем
type HealthType string

const (
	UnitType     HealthType = "unit"     // Юнит
	BuildingType HealthType = "building" // Здание
	HeroType     HealthType = "hero"     // Герой
)

// Health представляет компонент здоровья
type Health struct {
	Type     HealthType // Тип сущности (юнит или здание)
	MaxValue int        // Максимальное значение
	Value    int        // Текущее значение
}

// NewHealth создает новый компонент здоровья
func NewHealth(dType HealthType, maxValue int) (*Health, error) {
	validTypes := map[HealthType]bool{
		UnitType:     true,
		BuildingType: true,
		HeroType:     true,
	}
	if _, ok := validTypes[dType]; !ok {
		return nil, errors.ErrInvalidType
	}

	return &Health{
		Type:     dType,
		MaxValue: maxValue,
		Value:    maxValue,
	}, nil
}

// ApplyDamage уменьшает текущее HP на заданное значение
func (d *Health) ApplyDamage(damage int) {
	d.Value -= damage
	if d.Value < 0 {
		d.Value = 0
	}
}

// Heal увеличивает текущее HP на заданное значение
func (d *Health) Heal(amount int) {
	d.Value += amount
	if d.Value > d.MaxValue {
		d.Value = d.MaxValue
	}
}

// IsDead проверяет, мертва ли сущность (HP = 0)
func (d *Health) IsDead() bool {
	return d.Value == 0
}

// GetType возвращает тип HP
func (d *Health) GetType() HealthType {
	return d.Type
}

// Name - возвращает компонент
func (d *Health) Name() string {
	return "Health"
}
