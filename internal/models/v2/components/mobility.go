package components

import (
	"game-state/internal/models/v2/errors"
)

// Position представляет координаты в гексагональной системе
type Position struct {
	Q int // Координата Q
	R int // Координата R
	S int // Координата S
}

// Mobility представляет компонент, отвечающий за перемещение и видимость
type Mobility struct {
	Position Position // Текущая позиция объекта
	Speed    int      // Скорость перемещения
	Vision   int      // Радиус видимости
}

// NewMobility создает новый компонент перемещения
func NewMobility(q, r, s, speed, vision int) (*Mobility, error) {
	// Проверка на корректность гексагональных координат
	if q+r+s != 0 {
		return nil, errors.ErrInvalidCoordinates
	}

	return &Mobility{
		Position: Position{Q: q, R: r, S: s},
		Speed:    speed,
		Vision:   vision,
	}, nil
}

// Move перемещает объект на указанные координаты (с учетом скорости)
func (m *Mobility) Move(deltaQ, deltaR, deltaS int) error {
	newQ := m.Position.Q + deltaQ
	newR := m.Position.R + deltaR
	newS := m.Position.S + deltaS

	// Проверка на корректность новых координат
	if newQ+newR+newS != 0 {
		return errors.ErrInvalidCoordinates
	}

	m.Position.Q = newQ
	m.Position.R = newR
	m.Position.S = newS

	return nil
}

// Name возвращает имя компонента
func (m *Mobility) Name() string {
	return "mobility"
}
