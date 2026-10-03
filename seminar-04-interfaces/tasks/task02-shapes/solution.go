package shapes

import "math"

// Area возвращает площадь круга.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Area возвращает площадь прямоугольника.
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Area возвращает площадь квадрата.
func (s Square) Area() float64 {
	return s.Side * s.Side
}

// TotalArea возвращает суммарную площадь фигур из списка.
//
// Фигуры в списке могут быть любых типов, объявленных в types.go, в любом
// порядке. Пустой и nil-список дают 0.
func TotalArea(shapes []Shape) (sum float64) {
	if len(shapes) == 0 {
		return 0
	}

	for _, shape := range shapes {
		sum += shape.Area()
	}

	return
}

// MaxShape возвращает фигуру с наибольшей площадью из списка.
//
// Если площади равны, возвращается первая из фигур с этой площадью.
// Для пустого и nil-списка возвращает nil.
func MaxShape(shapes []Shape) (mx Shape) {
	if len(shapes) == 0 {
		return nil
	}

	mx = shapes[0]
	for _, shape := range shapes[1:] {
		if shape.Area() > mx.Area() {
			mx = shape
		}
	}

	return
}
