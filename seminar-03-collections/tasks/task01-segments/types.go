package segments

// Этот файл менять не нужно — тесты рассчитаны именно на эти объявления.

import "errors"

// ErrInvalidSize обозначает, что запрошенный размер куска недопустим.
var ErrInvalidSize = errors.New("invalid chunk size")

// ErrRange обозначает, что запрошенный диапазон выходит за границы или задан некорректно.
var ErrRange = errors.New("invalid range")
