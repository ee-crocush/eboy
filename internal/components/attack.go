package components

// Attack - компонент атаки для сущности.
type Attack struct {
	Damage  int  `json:"damage"`   // Урон
	IsRange bool `json:"is_range"` // Флаг дальнобойности
	Range   byte `json:"range"`    // Дальность атаки (в гексах)
}
