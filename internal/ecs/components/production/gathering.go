package production

// Gathering - компонент добычи ресурсов
type Gathering struct {
	Resource     Resource `json:"resource"`     // Ресурс, который добывается
	Productivity int      `json:"productivity"` // Сколько продукта добывается в минуту
}
