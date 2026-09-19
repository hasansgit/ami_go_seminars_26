package divmod

// DivMod возвращает частное и остаток от деления dividend на divisor.
//
// Если divisor не равен нулю, результат должен совпадать с тем, что дают
// выражения dividend / divisor и dividend % divisor в Go. Ошибка в этом
// случае равна nil.
//
// Если divisor == 0, нужно вернуть нулевые частное и остаток, а в качестве
// ошибки — ErrDivisionByZero.
//
// Паниковать при ожидаемой ошибке нельзя.
func DivMod(dividend, divisor int) (quotient, remainder int, err error) {
	panic("TODO: реализуйте DivMod")
}
