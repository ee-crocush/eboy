package production

import "time"

type ResourceCost struct {
	Resource Resource `json:"resource"` // Ресурс
	Amount   int      `json:"amount"`   // Количество
}

// ProductionComponent - интерфейс для производимого компонента
type ProductionComponent interface{}

// Production - компонент производства
type Production struct {
	Duration  time.Duration       `json:"duration"`  // Время, необходимое для производства            `json:"productivity"` // Сколько продукта добывается в минуту
	Resources []ResourceCost      `json:"resources"` // Список ресурсов, необходимые для производства
	Product   ProductionComponent `json:"product"`   // Компонент, который будет производиться
}
