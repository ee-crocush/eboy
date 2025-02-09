package errors

import "errors"

// Общие ошибки
var (
	ErrInvalidType = errors.New("invalid type") // Неверный тип
)

// Ошибки по координатам
var (
	ErrInvalidCoordinates = errors.New("invalid hex coordinates: Q + R + S must equal 0") // Неверные координаты
)

// Ошибки по компонентам
var (
	ErrComponentNotFound      = errors.New("component not found")
	ErrComponentAlreadyExists = errors.New("component already exists")
	ErrZeroDamage             = errors.New("damage must be greater than 0")
	ErrCooldownNegative       = errors.New("cooldown cannot be negative")
	ErrCooldown               = errors.New("attack is on cooldown")
)
