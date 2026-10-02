// Интерфейсы: неявное удовлетворение и fmt.Stringer.
//
//	go run ./seminar-04-interfaces/examples/02-interfaces
package main

import "fmt"

// Интерфейс — это набор сигнатур методов. Никакого implements:
// тип удовлетворяет интерфейсу сам, просто объявив нужные методы.
//
// Так устроен fmt.Stringer из стандартной библиотеки:
//
//	type Stringer interface {
//		String() string
//	}

// Percent — проценты, число от 0 до 100.
type Percent int

// String делает Percent строкером. fmt сам вызовет этот метод.
func (p Percent) String() string {
	return fmt.Sprintf("%d%%", int(p))
}

// Temperature — температура в градусах Цельсия.
type Temperature float64

func (t Temperature) String() string {
	return fmt.Sprintf("%.1f°C", float64(t))
}

// describe принимает интерфейс, а не конкретный тип: подойдёт любое
// значение, у которого есть String(). Это и есть полиморфизм в Go.
func describe(name string, s fmt.Stringer) {
	fmt.Printf("%-12s -> %s\n", name, s)
}

func main() {
	// Без метода String печать выглядела бы как "42" и "36.6".
	fmt.Println(Percent(42))       // 42%
	fmt.Println(Temperature(36.6)) // 36.6°C

	describe("загрузка", Percent(87))
	describe("на улице", Temperature(-3.5))

	// Интерфейс — тоже тип. Можно хранить разные типы в одном слайсе,
	// если все они удовлетворяют интерфейсу.
	values := []fmt.Stringer{Percent(100), Temperature(0)}
	for _, v := range values {
		fmt.Println(v)
	}
}
