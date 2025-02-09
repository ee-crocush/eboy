package components

// Mobility - компонент мобильности для сущности.
type Mobility struct {
	Speed  float32 `json:"speed"`  // Скорость передвижения
	Vision byte    `json:"vision"` // Дальность видимости (в гексах)
}
