// any, type assertion и type switch.
//
//	go run ./seminar-04-interfaces/examples/04-type-switch
package main

import "fmt"

// any — пустой интерфейс interface{}: ему удовлетворяет любой тип,
// потому что в нём ноль методов. Значение any хранит пару
// (конкретный тип, значение), и тип можно выяснить обратно.

// Type assertion: v.(T) — «достань из интерфейса значение типа T».
// Форма с двумя результатами (comma-ok) не паникует при промахе.
func describe(v any) string {
	if s, ok := v.(string); ok {
		return "строка «" + s + "»"
	}
	if n, ok := v.(int); ok {
		return fmt.Sprintf("число %d", n)
	}
	return "что-то ещё"
}

// Лестница из if с assertion — признак того, что нужен type switch:
// одна конструкция вместо серии проверок, с обязательной веткой default.
func classify(v any) string {
	switch x := v.(type) {
	case nil: // nil — тоже отдельный случай
		return "nil"
	case int:
		return fmt.Sprintf("int: %d", x)
	case string:
		return fmt.Sprintf("string длины %d", len(x))
	case bool:
		return fmt.Sprintf("bool: %t", x)
	case fmt.Stringer: // в case может стоять и интерфейс
		return "умеет String(): " + x.String()
	default:
		return fmt.Sprintf("неизвестный тип %T", x)
	}
}

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%.1f°C", c) }

func main() {
	fmt.Println(describe("го"))
	fmt.Println(describe(42))
	fmt.Println(describe(3.14))

	fmt.Println(classify(nil))
	fmt.Println(classify(7))
	fmt.Println(classify("семинар"))
	fmt.Println(classify(true))
	fmt.Println(classify(Celsius(21.5)))
	fmt.Println(classify([]int{1, 2, 3}))

	// Ловушка: assert v.(T) без ok при промахе ПАНИКУЕТ.
	// Раскомментируйте и запустите:
	//
	// var v any = "строка"
	// _ = v.(int) // panic: interface conversion
}
