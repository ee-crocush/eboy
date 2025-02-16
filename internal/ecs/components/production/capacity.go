package production

// Capacity - характеристики емкости для хранения ресурсов.
type Capacity struct {
	Capacity    int `json:"capacity"`     // Сколько всего ресурсов хранит объект
	CapacityNow int `json:"capacity_now"` // Текущая емкость
}
