// Этот файл менять не нужно — тесты рассчитаны именно на эти объявления.

package shapes

// Shape — геометрическая фигура, умеющая сообщать свою площадь.
//
// Это единственный интерфейс задачи: реализации ниже удовлетворяют ему
// неявно, достаточно объявить метод Area() float64.
type Shape interface {
	Area() float64
}

// Circle — круг заданного радиуса.
type Circle struct {
	Radius float64
}

// Rectangle — прямоугольник с заданными сторонами.
type Rectangle struct {
	Width  float64
	Height float64
}

// Square — квадрат с заданной стороной.
type Square struct {
	Side float64
}
