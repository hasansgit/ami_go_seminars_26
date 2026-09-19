// Структуры, указатели и инкапсуляция на уровне пакета.
//
// Запуск:
//
//	go run ./seminar-02-language/examples/06-structs-pointers
//
// Аналогия для тех, кто знает C++: структура — это struct без наследования
// и без конструктора, а указатель — это ссылка, которую нельзя сдвинуть
// арифметикой. Аналогия для тех, кто знает Python: это класс с полями,
// но без __init__ и без private-по-соглашению — приватность тут настоящая.
package main

import "fmt"

// Point — обычная структура: набор именованных полей.
// Имена с большой буквы видны из других пакетов, с маленькой — нет.
type Point struct {
	X, Y int
}

// counter прячет своё поле: value с маленькой буквы, поэтому снаружи
// пакета его не прочитать и не записать. Это и есть вся инкапсуляция в Go —
// никаких ключевых слов private/public, решает регистр первой буквы.
type counter struct {
	value int
}

func main() {
	zeroValues()
	valueVsPointer()
	encapsulation()
}

func zeroValues() {
	fmt.Println("--- нулевое значение структуры ---")

	// Структура инициализируется автоматически, как и любая переменная:
	// каждое поле получает своё нулевое значение. Конструктор не нужен.
	var p Point
	fmt.Printf("  var p Point         -> %+v\n", p)

	// Три способа создать структуру со значениями.
	byFields := Point{X: 3, Y: 4} // по именам полей — так пишут почти всегда
	byOrder := Point{3, 4}        // по порядку — только для очень коротких структур
	byPointer := &Point{X: 3}     // сразу указатель, Y останется нулевым

	fmt.Printf("  Point{X: 3, Y: 4}   -> %+v\n", byFields)
	fmt.Printf("  Point{3, 4}         -> %+v\n", byOrder)
	fmt.Printf("  &Point{X: 3}        -> %+v\n", byPointer)

	// Структуры сравнимы через ==, если сравнимы все их поля.
	fmt.Println("  byFields == byOrder:", byFields == byOrder)
}

func valueVsPointer() {
	fmt.Println("--- значение против указателя ---")

	p := Point{X: 1, Y: 1}

	moveByValue(p)
	fmt.Printf("  после moveByValue   -> %+v  (ничего не изменилось)\n", p)

	moveByPointer(&p)
	fmt.Printf("  после moveByPointer -> %+v  (изменилось)\n", p)

	// ЛОВУШКА: у nil-указателя нет полей, обращение к ним — паника.
	// Поэтому функции, принимающие указатель, обычно проверяют его на nil.
	var missing *Point
	fmt.Println("  missing == nil:", missing == nil)
}

// moveByValue получает КОПИЮ структуры. Оригинал вызывающей стороны
// такая функция изменить не может — в Go всё передаётся по значению.
func moveByValue(p Point) {
	p.X += 10
}

// moveByPointer получает адрес и потому меняет оригинал.
// Обратите внимание: писать (*p).X не нужно, Go разыменовывает сам.
func moveByPointer(p *Point) {
	p.X += 10
}

func encapsulation() {
	fmt.Println("--- инкапсуляция ---")

	// Внутри своего пакета приватное поле доступно — ограничение работает
	// на границе пакета, а не на границе типа.
	c := &counter{}
	increment(c)
	increment(c)

	fmt.Println("  значение счётчика:", currentValue(c))
	fmt.Println("  из другого пакета c.value было бы недоступно —")
	fmt.Println("  остались бы только функции increment и currentValue")
}

func increment(c *counter) {
	c.value++
}

func currentValue(c *counter) int {
	if c == nil {
		return 0
	}
	return c.value
}
