package money

import (
	"fmt"
)

// String возвращает строковое представление суммы в формате
// "рубли.копейки ВАЛЮТА".
//
// Копейки занимают ровно два знака после точки: Money{Amount: 5,
// Currency: "RUB"} — это "0.05 RUB". Сумма, кратная рублю, всё равно
// печатает копейки: Money{Amount: 1200, Currency: "RUB"} — это
// "12.00 RUB".
//
// Отрицательная сумма получает один знак минус перед числом:
// Money{Amount: -5, Currency: "RUB"} — это "-0.05 RUB".
func (m Money) String() (res string) {
	return fmt.Sprintf("%.2f %v", float32(m.Amount)/100, m.Currency)
}
