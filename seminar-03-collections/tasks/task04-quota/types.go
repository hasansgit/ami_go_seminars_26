// Этот файл менять не нужно — тесты рассчитаны именно на эти объявления.

package quota

import "errors"

// Op — одна операция расхода: пользователь user тратит amount единиц квоты.
type Op struct {
	User   string
	Amount int
}

// Usage — суммарный расход одного пользователя.
type Usage struct {
	User  string
	Total int
}

// ErrUnknownUser обозначает, что операция ссылается на пользователя без лимита.
var ErrUnknownUser = errors.New("unknown user")

// ErrInvalidAmount обозначает, что величина операции недопустима.
var ErrInvalidAmount = errors.New("invalid amount")

// ErrOverQuota обозначает, что операция превышает оставшийся лимит пользователя.
var ErrOverQuota = errors.New("quota exceeded")
