package base

// Position - компонент позиции в гекосагональной сетке
type Position struct {
	Q int `json:"q"` // Координата Q
	R int `json:"r"` // Координата R
	S int `json:"s"` // Координата S
}
