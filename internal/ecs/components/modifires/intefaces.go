package modifires

// Modifier - модификатор параметров объектов
type Modifier interface {
	ApplyTo(properties any) any // Применить модификацию
}
