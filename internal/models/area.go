//area.go - модели игровых объектов каждой игровой зоны
//У игрока может быть несколько игровых зон, и на каждой будет свои объекты

package models

import (
	"errors"
	"github.com/shopspring/decimal"
)

var ErrNoneZeroRadius = errors.New("радиус игровой зоны должен быть больше 0")

// Coordinates - Гексогональная сетка
type Coordinates struct {
	Q int8 `json:"q"`
	R int8 `json:"r"`
	S int8 `json:"s"`
}

// Area - игровая зона
type Area struct {
	Id     int32 `db:"id"`      // ID Объекта
	UserId int32 `db:"user_id"` // ID пользователя
	Radius int8  `db:"radius"`  // Радиус зоны
}

// Validate проверяет, что радиус зоны не нулевой
func (a *Area) Validate() error {
	if a.Radius < 1 {
		return ErrNoneZeroRadius
	}
	return nil
}

// AreaNeutral - нейтральный объект на игровой зоне
type AreaNeutral struct {
	Id          int32           `db:"id"`          // ID
	AreaId      int32           `db:"area_id"`     // ID игровой зоны
	NeutralId   int32           `db:"neutral_id"`  // ID нейтрального объекта
	Capacity    decimal.Decimal `db:"capacity"`    // Емкость текущая
	Coordinates Coordinates     `db:"coordinates"` // Координаты
}

// AreaBuilding - здание на игровой зоне
type AreaBuilding struct {
	Id              int32                   `db:"id"`                                     // ID
	AreaId          int32                   `db:"area_id"`                                // ID игровой зоны
	BuildingId      int32                   `db:"building_id"`                            // ID здания
	Characteristics BuildingCharacteristics `db:"characteristics" json:"characteristics"` // Характеристики текущие
	Coordinates     Coordinates             `db:"coordinates"`                            // Координаты
}

// AreaHero - герои на игровой зоне
type AreaHero struct {
	Id              int32               `db:"id"`                                     // ID
	AreaId          int32               `db:"area_id"`                                // ID игровой зоны
	HeroId          int32               `db:"hero_id"`                                // ID героя
	Characteristics HeroCharacteristics `db:"characteristics" json:"characteristics"` // Характеристики текущие
	Experience      decimal.Decimal     `db:"experience"`                             // Опыт текущий
	Coordinates     Coordinates         `db:"coordinates"`                            // Координаты
}

// AreaUnit - юнит на игровой зоне
type AreaUnit struct {
	Id              int32               `db:"id"`                                     // ID
	AreaId          int32               `db:"area_id"`                                // ID игровой зоны
	UnitId          int32               `db:"unit_id"`                                // ID юнита
	Characteristics UnitCharacteristics `db:"characteristics" json:"characteristics"` // Характеристики текущие
	Experience      decimal.Decimal     `db:"experience"`                             // Опыт текущий
	Coordinates     Coordinates         `db:"coordinates"`                            // Координаты
}

// AreaEnemy - враг на игровой зоне
type AreaEnemy struct {
	Id              int32                `db:"id"`                                     // ID
	AreaId          int32                `db:"area_id"`                                // ID игровой зоны
	EnemyId         int32                `db:"enemy_id"`                               // ID врага
	Characteristics EnemyCharacteristics `db:"characteristics" json:"characteristics"` // Характеристики текущие
	Coordinates     Coordinates          `db:"coordinates"`                            // Координаты
}
