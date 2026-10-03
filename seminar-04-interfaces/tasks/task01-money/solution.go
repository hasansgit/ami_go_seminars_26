package money

import "strconv"

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
	if m.Amount < 0 {
		res = "-"
		m.Amount *= -1
	}

	res += strconv.Itoa(m.Amount/100) + "."

	if m.Amount%100 < 10 {
		res += "0"
	}
	res += strconv.Itoa(m.Amount%100) + " " + m.Currency

	return
}
