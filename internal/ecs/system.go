package ecs

// System - интерфейс для всех систем.
type System interface {
	Update(entities []*Entity) // Обновление всех сущностей в системе
}
