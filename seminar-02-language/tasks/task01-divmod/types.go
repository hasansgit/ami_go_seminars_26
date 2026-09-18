// Этот файл менять не нужно — тесты рассчитывают именно на эти объявления.

package divmod

import "errors"

// ErrDivisionByZero сообщает о попытке деления на ноль.
var ErrDivisionByZero = errors.New("division by zero")
