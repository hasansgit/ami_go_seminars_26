// Этот файл менять не нужно — тесты рассчитывают именно на эти объявления.

package permissions

// Permission представляет набор независимых прав доступа.
type Permission uint8

const (
	Read Permission = 1 << iota
	Write
	Execute
)
