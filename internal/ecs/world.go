package ecs

// World - мир, который содержит все сущности и системы.
type World struct {
	entities []*Entity
	systems  []System // Системы, которые обновляют мир
}

// NewWorld - создает новый мир.
func NewWorld() *World {
	return &World{}
}

// AddEntity - добавляет сущность в мир.
func (w *World) AddEntity(entity *Entity) {
	w.entities = append(w.entities, entity)
}

// AddSystem - добавляет систему в мир.
func (w *World) AddSystem(system System) {
	w.systems = append(w.systems, system)
}

// Update - обновляет все системы в мире.
func (w *World) Update() {
	for _, system := range w.systems {
		system.Update(w.entities)
	}
}
