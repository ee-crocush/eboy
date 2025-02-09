//world_objects.go - модели игровых объектов.
//Данные модели в любом случае нужны, так как не будем же
//рандомно создать какие то объекты, нужна база всех возможных объектов,
//которые будут доступны пользователю

package old

import (
	"errors"
	"github.com/shopspring/decimal"
	"time"
)

// ErrNoneZeroLevel - уровень должнен быть больше 0, скорее всего даже избыточен, так как
// уровень всегда больше 0 и относится к базовым объектам
var ErrNoneZeroLevel = errors.New("уровень должнен быть больше 0")

// ObjectType определяет доступные значения для поля Type.
type ObjectType string

const (
	ObjectTypeNeutral  ObjectType = "neutrals"  // Нейтральные объекты
	ObjectTypeBuilding ObjectType = "buildings" // Здания
	ObjectTypeHero     ObjectType = "heroes"    // Герои
	ObjectTypeUnit     ObjectType = "units"     // Юниты
	ObjectTypeEnemy    ObjectType = "enemies"   // Враги
)

// Object - игровой объект
type Object struct {
	Id   int32      `db:"id"`   // ID Объекта
	Type ObjectType `db:"type"` // Тип объекта
}

// Validate проверяет, что тип объекта валиден.
func (o *Object) Validate() error {
	switch o.Type {
	case ObjectTypeNeutral, ObjectTypeBuilding, ObjectTypeHero, ObjectTypeUnit, ObjectTypeEnemy:
		return nil
	default:
		return errors.New("неверный тип объекта: " + string(o.Type))
	}
}

// Neutral - нейтральный объект (Деревья, реки и т.д.)
type Neutral struct {
	id          int32           `db:"id"`           // ID
	ObjectId    int32           `db:"object_id"`    // ID объекта
	Name        string          `db:"name"`         // Наименование
	ProductId   int32           `db:"product_id"`   // ID ресурса (если это гора, то ресурс может не будет добываться
	ProdCof     decimal.Decimal `db:"prod_cof"`     // Коэффициент производительности
	Capacity    decimal.Decimal `db:"capacity"`     // Емкость
	CanGathered bool            `db:"can_gathered"` // флаг, обозначающий, что объект может быть собран
}

// BaseCharacteristics - базовые характеристики (хп и броня)
type BaseCharacteristics struct {
	HP    int `json:"hp"`    // Здоровье
	Armor int `json:"armor"` // Броня
}

// BuildingCharacteristics - характеристики здания
type BuildingCharacteristics struct {
	BaseCharacteristics
}

// Building - здание
type Building struct {
	id              int32                   `db:"id"`                                     // ID
	ObjectId        int32                   `db:"object_id"`                              // ID объекта
	Name            string                  `db:"name"`                                   // Наименование
	Characteristics BuildingCharacteristics `db:"characteristics" json:"characteristics"` // Характеристики базовые
	Level           int32                   `db:"level"`                                  // Уровень здания
	UpgradePrice    decimal.Decimal         `db:"upgrade_price"`                          // Стоимость улучшения
}

// Validate проверяет, уровень здание не ноль
func (b *Building) Validate() error {
	if b.Level < 1 {
		return ErrNoneZeroLevel
	}
	return nil
}

// HeroCharacteristics представляет характеристики героя. (дубль с характеристиками юнита)
type HeroCharacteristics struct {
	BaseCharacteristics
	Speed       decimal.Decimal `json:"speed"`        // Скорость перемещения героя
	Vision      int             `json:"vision"`       // Дальность обзора героя
	IsRange     bool            `json:"range"`        // Флаг, определяющий, является ли герой дальнобойным
	AttackRange decimal.Decimal `json:"attack_range"` // Дальность атаки героя
	Damage      int             `json:"damage"`       // Текущий урон героя
}

// Hero - герои
type Hero struct {
	id              int32               `db:"id"`                                     // ID
	ObjectId        int32               `db:"object_id"`                              // ID объекта
	Name            string              `db:"name"`                                   // Наименование
	Characteristics HeroCharacteristics `db:"characteristics" json:"characteristics"` // Характеристики базовые
	ExperienceToUp  decimal.Decimal     `db:"experience_to_up"`                       // Опыт до следующего улучшения
	Level           int32               `db:"level"`                                  // Уровень
}

// Validate проверяет, уровень героя не ноль
func (h *Hero) Validate() error {
	if h.Level < 1 {
		return ErrNoneZeroLevel
	}
	return nil
}

// AbilityCharacteristics представляет характеристики способности.
type AbilityCharacteristics struct {
	IsPassive bool            `json:"is_passive"` // Флаг, определяющий, является ли способность пассивной
	Radius    decimal.Decimal `json:"radius"`     // Радиус действия способности
	Cooldown  time.Duration   `json:"cooldown"`   // Время перезарядки способности
	Damage    decimal.Decimal `json:"damage"`     // Урон от способности
}

// Ability - способность
type Ability struct {
	id              int32                  `db:"id"`                                     // ID
	Name            string                 `db:"name"`                                   // Наименование
	Characteristics AbilityCharacteristics `db:"characteristics" json:"characteristics"` // Характеристики
	Level           int32                  `db:"level"`                                  // Уровень
}

// Validate проверяет, уровень способности не ноль
func (a *Ability) Validate() error {
	if a.Level < 1 {
		return ErrNoneZeroLevel
	}
	return nil
}

// HeroAbility - Способности героев, нужна ли модель?
type HeroAbility struct {
	HeroId    int32 `db:"hero_id"`    // ID героя
	AbilityId int32 `db:"ability_id"` // Способность
}

// UnitCharacteristics представляет характеристики юнита. (дубль с характеристиками героя)
type UnitCharacteristics struct {
	HeroCharacteristics
}

// Unit - юнит
type Unit struct {
	Id              int32               `db:"id"`                                     // ID
	ObjectId        int32               `db:"object_id"`                              // ID объекта
	Name            string              `db:"name"`                                   // Наименование
	Characteristics UnitCharacteristics `db:"characteristics" json:"characteristics"` // Характеристики базовые
	ExperienceToUp  decimal.Decimal     `db:"experience_to_up"`                       // Опыт до следующего улучшения
	Level           int32               `db:"level"`                                  // Уровень
}

// Validate проверяет, уровень юнита не ноль
func (u *Unit) Validate() error {
	if u.Level < 1 {
		return ErrNoneZeroLevel
	}
	return nil
}

// EnemyCharacteristics представляет характеристики врага. (дубль с характеристиками героя)
type EnemyCharacteristics struct {
	HeroCharacteristics
	Experience decimal.Decimal `json:"experience"` // Получаемый опыт за убииство
}

// Enemy - враг
type Enemy struct {
	Id              int32                `db:"id"`                                     // ID
	ObjectId        int32                `db:"object_id"`                              // ID объекта
	Name            string               `db:"name"`                                   // Наименование
	Characteristics EnemyCharacteristics `db:"characteristics" json:"characteristics"` // Характеристики базовые
	Level           int32                `db:"level"`                                  // Уровень
}

// Validate проверяет, уровень врага не ноль
func (e *Enemy) Validate() error {
	if e.Level < 1 {
		return ErrNoneZeroLevel
	}
	return nil
}
