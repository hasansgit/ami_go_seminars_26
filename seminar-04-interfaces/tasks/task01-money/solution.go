package money

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
func (m Money) String() string {
	panic("TODO: реализуйте Money.String")
}
