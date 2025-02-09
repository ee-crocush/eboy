// models/objects.go
package v1

import (
	"fmt"
)

var (
	ErrLowExp       = fmt.Errorf("not enough experience")
	ErrLowRes       = fmt.Errorf("not enough resources")
	ErrNotSatisfied = fmt.Errorf("requirements not satisfied")
)

// ObjectType - типы объектов в игре.
type ObjectType string

const (
	// Типы объектов
	TypeNeutral  ObjectType = "neutrals"  // Нейтральные объекты
	TypeBuilding ObjectType = "buildings" // Здания
	TypeHero     ObjectType = "heroes"    // Герои
	TypeUnit     ObjectType = "units"     // Юниты
	TypeEnemy    ObjectType = "enemies"   // Враги
)

// Resource - типы ресурсов в игре.
type Resource string

const (
	// Доступные ресурсы
	Gold  Resource = "gold"  // Золото
	Water Resource = "water" // Вода
	Woods Resource = "woods" // Дерево
)

// Coordinates - координаты объекта в гексагональной сетке.
type Coordinates struct {
	Q int8 `json:"q"`
	R int8 `json:"r"`
	S int8 `json:"s"`
}

// ObjectProperties - базовые параметры объекта (ID, тип и координаты).
type ObjectProperties struct {
	Id          int64       `json:"id"`          // Идентификатор объекта
	Type        ObjectType  `json:"type"`        // Тип объекта
	Coordinates Coordinates `json:"coordinates"` // Координаты
}

// CapacityProperties - характеристики емкости для хранения ресурсов.
type CapacityProperties struct {
	Capacity    int `json:"capacity"`     // Сколько всего ресурсов хранит объект
	CapacityNow int `json:"capacity_now"` // Текущая емкость
}

// ProductionProperties - характеристики для объектов, участвующих в производстве.
type ProductionProperties struct {
	Resource     Resource `json:"resource"`     // Ресурс, который добывается
	Productivity int      `json:"productivity"` // Сколько продукта добывается в минуту
}

// HealthProperties - характеристики прочности (здоровье, броня).
type HealthProperties struct {
	HP    int `json:"hp"`     // Полное здоровье объекта
	HPNow int `json:"hp_now"` // Текущее здоровье
	Armor int `json:"armor"`  // Броня
}

// MobilityProperties - характеристики мобильности (скорость и дальность видимости).
type MobilityProperties struct {
	Speed  float32 `json:"speed"`  // Скорость передвижения объекта
	Vision byte    `json:"vision"` // Дальность видимости (в гексах)
}

// AttackProperties - характеристики атаки (урон, дальность, возможность дальнего боя).
type AttackProperties struct {
	Damage  int  `json:"damage"`   // Урон
	IsRange bool `json:"is_range"` // Флаг дальнобойности
	Range   byte `json:"range"`    // Дальность атаки (в гексах)
}

// Unit - объект юнита с его характеристиками.
type Unit struct {
	ObjectProperties
	Health    *HealthProperties   `json:"Health"`    // Прочность
	Mobility  *MobilityProperties `json:"mobility"`  // Мобильность
	Attack    *AttackProperties   `json:"attack"`    // Атака
	Modifiers []Modifier          `json:"modifiers"` // Модификаторы
}

// ApplyModifiers - применяет все модификаторы к юниту.
func (obj *Unit) ApplyModifiers() {
	for _, modifier := range obj.Modifiers {
		if obj.Health != nil {
			obj.Health = modifier.ApplyTo(obj.Health).(*HealthProperties)
		}
		if obj.Mobility != nil {
			obj.Mobility = modifier.ApplyTo(obj.Mobility).(*MobilityProperties)
		}
		if obj.Attack != nil {
			obj.Attack = modifier.ApplyTo(obj.Attack).(*AttackProperties)
		}
	}
}
