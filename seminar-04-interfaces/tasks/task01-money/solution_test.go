package money

import (
	"fmt"
	"testing"
)

func TestString(t *testing.T) {
	testCases := []struct {
		name string
		m    Money
		want string
	}{
		{name: "только копейки", m: Money{Amount: 5, Currency: "RUB"}, want: "0.05 RUB"},
		{name: "рубли и копейки", m: Money{Amount: 1234, Currency: "RUB"}, want: "12.34 RUB"},
		{name: "сумма, кратная рублю, печатает копейки", m: Money{Amount: 1200, Currency: "RUB"}, want: "12.00 RUB"},
		{name: "ноль", m: Money{Amount: 0, Currency: "USD"}, want: "0.00 USD"},
		{name: "одна копейка", m: Money{Amount: 1, Currency: "RUB"}, want: "0.01 RUB"},
		{name: "десять копеек", m: Money{Amount: 10, Currency: "RUB"}, want: "0.10 RUB"},
		{name: "отрицательные копейки", m: Money{Amount: -5, Currency: "RUB"}, want: "-0.05 RUB"},
		{name: "отрицательные рубли и копейки", m: Money{Amount: -1234, Currency: "RUB"}, want: "-12.34 RUB"},
		{name: "большая сумма", m: Money{Amount: 100000000, Currency: "EUR"}, want: "1000000.00 EUR"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.m.String(); got != tc.want {
				t.Errorf("Money{Amount: %d, Currency: %q}.String() = %q, ожидалось %q",
					tc.m.Amount, tc.m.Currency, got, tc.want)
			}
		})
	}
}

// fmt.Sprint обязан вызвать String() сам — в этом и смысл интерфейса
// fmt.Stringer.
func TestStringIsUsedByFmt(t *testing.T) {
	testCases := []struct {
		name string
		m    Money
	}{
		{name: "положительная сумма", m: Money{Amount: 999, Currency: "RUB"}},
		{name: "отрицательная сумма", m: Money{Amount: -42, Currency: "USD"}},
		{name: "нулевая сумма", m: Money{Amount: 0, Currency: "EUR"}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := fmt.Sprint(tc.m); got != tc.m.String() {
				t.Errorf("fmt.Sprint(%+v) = %q, ожидалось совпадение с String(): %q",
					tc.m, got, tc.m.String())
			}
		})
	}
}

// Значение Money, а не указатель на него, обязано удовлетворять
// fmt.Stringer: суммы передаются и хранятся по значению.
func TestValueSatisfiesStringer(t *testing.T) {
	var m Money = Money{Amount: 100, Currency: "RUB"}
	var _ fmt.Stringer = m
}
