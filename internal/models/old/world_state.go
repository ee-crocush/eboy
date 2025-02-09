//world_state.go - состояние мира игрока

package old

import (
	"github.com/shopspring/decimal"
	"time"
)

// ActionType определяет доступные действия
type ActionType int8

const (
	ActionMove    ActionType = 1 // MOVE - движение
	ActionAttack  ActionType = 2 // ATTACK - атака
	ActionBuild   ActionType = 3 // BUILD - строительство
	ActionUpgrade ActionType = 4 // UPGRADE - улучшение
)

// StatusType определяет доступные статусы действия
type StatusType string

const (
	StatusDone     StatusType = "DONE"     // DONE - выполнено
	StatusProcess  StatusType = "PROCESS"  // PROCESS - в процессе
	StatusCanceled StatusType = "CANCELED" // CANCELED - отменено
)

// League - модель лиги
type League struct {
	Id        int32  `db:"id"`        // Идентификатор лиги
	Name      string `db:"name"`      // Название лиги
	Authority int    `db:"authority"` // Уровень авторитета лиги
}

// User - модель пользователя
type User struct {
	UserId   int32           `db:"user_id"`   // Во всех сервисах оперируем user_id
	LeagueId string          `db:"league_id"` // Лига
	Balance  decimal.Decimal `db:"balance"`   // Баланс пользователя
	Level    int             `db:"level"`     // Уровень пользователя
}

// Resource - модель ресурса
type Resource struct {
	Id   int32  `db:"id"`   // Во всех сервисах оперируем user_id
	Name string `db:"name"` // Наименование ресурса
}

// UserResource - Ресурсы пользователя
type UserResource struct {
	UserId     int32           `db:"user_id"`     // Пользователь
	ResourceId int32           `db:"resource_id"` // Ресурс
	Value      decimal.Decimal `db:"value"`       // Значение ресурса
}

// Action - модель действия пользователя
type Action struct {
	Id             int32         `db:"id"`               // ID действия
	UserId         int32         `db:"user_id"`          // Пользователь
	AreaId         int32         `db:"area_id"`          // Игровая зона
	ObjectSourceId int32         `db:"object_source_id"` // ID источника действия (не обязательно внешний ключ, т.к. может ссылаться на разные типы)
	SourceType     ObjectType    `db:"source_type"`      // Тип источника действияObjectSourceId int32      `db:"object_source_id"` // ID источника действия (не обязательно внешний ключ, т.к. может ссылаться на разные типы)
	ObjectDestId   int32         `db:"object_dest_id"`   // ID цели действия
	DestType       ObjectType    `db:"source_type"`      // Тип цели действия
	ActionType     ActionType    `db:"action_type"`      // Тип действия
	StartTime      time.Time     `db:"start_time"`       // Время начала действия
	Duration       time.Duration `db:"duration"`         // Продолжительность действия
	StatusType     StatusType    `db:"status_type"`      // Статус действия
}

// WorldState - состояние мира игрока
type WorldState struct {
	SessionId int32     `db:"session_id"` // ID сессии
	UserId    int32     `db:"user_id"`    // ID пользователя
	AreaId    int32     `db:"area_id"`    // ID игровой зоны
	ActionId  int32     `db:"action_id"`  // ID действия
	TimeStamp time.Time `db:"time_stamp"` // Время состояния
}
