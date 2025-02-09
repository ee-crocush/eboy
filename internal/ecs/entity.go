package ecs

// Entity - базовая структура сущности.
type Entity struct {
	ID         int64
	components map[string]interface{} // Компоненты, привязанные к сущности
}

// NewEntity - создает новую сущность с уникальным ID.
func NewEntity(id int64) *Entity {
	return &Entity{
		ID:         id,
		components: make(map[string]interface{}),
	}
}

// AddComponent - добавляет компонент к сущности.
func (e *Entity) AddComponent(name string, component interface{}) {
	e.components[name] = component
}

// GetComponent - возвращает компонент по имени.
func (e *Entity) GetComponent(name string) interface{} {
	return e.components[name]
}
