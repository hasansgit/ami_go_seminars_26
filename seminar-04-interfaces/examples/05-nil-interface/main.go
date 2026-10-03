// Главная ловушка интерфейсов: интерфейс с nil-указателем внутри
// сам НЕ равен nil.
//
//	go run ./seminar-04-interfaces/examples/05-nil-interface
package main

import "fmt"

// Интерфейсное значение — это пара (конкретный тип, значение).
// Интерфейс равен nil, только когда ОБЕ части пусты.

type Config struct {
	Host string
}

func (c *Config) String() string {
	if c == nil {
		return "<no config>"
	}
	return c.Host
}

// findConfig возвращает *Config и ошибку... и вот тут ловушка:
// когда конфиг не найден, функция кладёт в error-переменную типа
// fmt.Stringer ТИПИЗИРОВАННЫЙ nil.
func defaultConfig() fmt.Stringer {
	var c *Config // c == nil, но у него есть тип *Config
	return c      // интерфейс = (тип *Config, nil) — это НЕ nil!
}

func main() {
	var empty fmt.Stringer                                // настоящий nil-интерфейс: ни типа, ни значения
	fmt.Println("пустой интерфейс == nil:", empty == nil) // true

	c := defaultConfig()
	fmt.Println("интерфейс с (*Config)(nil) == nil:", c == nil) // false!

	// Поэтому c нельзя проверять на nil — проверка не сработает,
	// а вызов метода уйдёт в nil-получатель. Хорошо, что наш String()
	// это предусмотрел:
	fmt.Println("вызов метода на nil-получателе:", c) // <no config>

	// Правила выживания:
	// 1. Функция, возвращающая интерфейс, в случае «пусто» возвращает
	//    голый nil, а не типизированный:
	//        var c *Config
	//        if c == nil { return nil } // return c — БАГ
	// 2. Методы, допускающие nil-получатель, проверяют его первой строкой
	//    (как String выше) — так устроены, например, методы protobuf.
}
