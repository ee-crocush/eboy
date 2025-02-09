package components

import (
	"game-state/internal/models/v2/errors"
	"time"
)

// Attack представляет компонент атаки
type Attack struct {
	Damage   int           // Наносимый урон
	Range    int           // Дальность атаки
	Cooldown time.Duration // Время перезарядки

	lastAttackTime time.Time // Время последней атаки
}

// NewAttack создает новый компонент атаки
func NewAttack(damage, attackRange int, cooldown time.Duration) (*Attack, error) {
	if damage <= 0 {
		return nil, errors.ErrZeroDamage
	}
	if cooldown < 0 {
		return nil, errors.ErrCooldownNegative
	}

	return &Attack{
		Damage:   damage,
		Range:    attackRange,
		Cooldown: cooldown,
	}, nil
}

// CanAttack проверяет, готов ли компонент к атаке
func (a *Attack) CanAttack() bool {
	return time.Since(a.lastAttackTime) >= a.Cooldown
}

// PerformAttack выполняет атаку, если это возможно
func (a *Attack) PerformAttack() error {
	if !a.CanAttack() {
		return errors.ErrCooldown
	}

	a.lastAttackTime = time.Now()
	return nil
}

// Name возвращает имя компонента
func (a *Attack) Name() string {
	return "attack"
}
