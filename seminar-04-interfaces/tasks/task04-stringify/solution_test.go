package stringify

import "testing"

// price — тип с собственным методом String, для проверки ветки fmt.Stringer.
type price struct {
	amount int
}

func (p price) String() string {
	return "price!"
}

func TestStringify(t *testing.T) {
	testCases := []struct {
		name string
		v    any
		want string
	}{
		{name: "nil", v: nil, want: "<nil>"},
		{name: "строка", v: "привет", want: "привет"},
		{name: "пустая строка", v: "", want: ""},
		{name: "истина", v: true, want: "true"},
		{name: "ложь", v: false, want: "false"},
		{name: "положительное int", v: 42, want: "42"},
		{name: "отрицательное int", v: -7, want: "-7"},
		{name: "ноль int", v: 0, want: "0"},
		{name: "float64 с дробной частью", v: 1.5, want: "1.50"},
		{name: "float64 без дробной части", v: 3.0, want: "3.00"},
		{name: "отрицательный float64", v: -0.126, want: "-0.13"},
		{name: "float64 округляется до двух знаков", v: 2.004, want: "2.00"},
		{name: "срез строк", v: []string{"a", "b", "c"}, want: "[a, b, c]"},
		{name: "срез из одной строки", v: []string{"solo"}, want: "[solo]"},
		{name: "пустой срез строк", v: []string{}, want: "[]"},
		{name: "nil-срез строк", v: []string(nil), want: "[]"},
		{name: "тип с методом String", v: price{amount: 100}, want: "price!"},
		{name: "неподдерживаемый тип: срез int", v: []int{1, 2}, want: "<unknown>"},
		{name: "неподдерживаемый тип: int64", v: int64(42), want: "<unknown>"},
		{name: "неподдерживаемый тип: структура", v: struct{ X int }{X: 1}, want: "<unknown>"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := Stringify(tc.v); got != tc.want {
				t.Errorf("Stringify(%#v) = %q, ожидалось %q", tc.v, got, tc.want)
			}
		})
	}
}

// Ветка fmt.Stringer обязана проверяться раньше конкретных типов: строковый
// алиас с методом String() идёт по String(), а не по ветке string.
type brand string

func (b brand) String() string { return "brand:" + string(b) }

func TestStringerWinsOverUnderlyingType(t *testing.T) {
	if got := Stringify(brand("go")); got != "brand:go" {
		t.Errorf("Stringify(brand) = %q, ожидалось %q", got, "brand:go")
	}
}
