package shapes

// Area возвращает площадь круга.
func (c Circle) Area() float64 {
	panic("TODO: реализуйте Circle.Area")
}

// Area возвращает площадь прямоугольника.
func (r Rectangle) Area() float64 {
	panic("TODO: реализуйте Rectangle.Area")
}

// Area возвращает площадь квадрата.
func (s Square) Area() float64 {
	panic("TODO: реализуйте Square.Area")
}

// TotalArea возвращает суммарную площадь фигур из списка.
//
// Фигуры в списке могут быть любых типов, объявленных в types.go, в любом
// порядке. Пустой и nil-список дают 0.
func TotalArea(shapes []Shape) float64 {
	panic("TODO: реализуйте TotalArea")
}

// MaxShape возвращает фигуру с наибольшей площадью из списка.
//
// Если площади равны, возвращается первая из фигур с этой площадью.
// Для пустого и nil-списка возвращает nil.
func MaxShape(shapes []Shape) Shape {
	panic("TODO: реализуйте MaxShape")
}
