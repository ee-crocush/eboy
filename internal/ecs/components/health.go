package components

// Health - компонент здоровья для сущности.
type Health struct {
	HP    int `json:"hp"`     // Полное здоровье
	HPNow int `json:"hp_now"` // Текущее здоровье
	Armor int `json:"armor"`  // Броня
}

// ApplyDamage - применяет урон к здоровью.
func (h *Health) ApplyDamage(damage int) {
	effectiveDamage := damage - h.Armor
	if effectiveDamage < 0 {
		effectiveDamage = 0
	}

	h.HPNow -= effectiveDamage
	if h.HPNow < 0 {
		h.HPNow = 0
	}
}
