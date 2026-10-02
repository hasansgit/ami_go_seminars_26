package shapes

import (
	"math"
	"testing"
)

const eps = 1e-9

func TestArea(t *testing.T) {
	testCases := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "круг единичного радиуса", shape: Circle{Radius: 1}, want: math.Pi},
		{name: "круг радиуса 2", shape: Circle{Radius: 2}, want: 4 * math.Pi},
		{name: "прямоугольник", shape: Rectangle{Width: 3, Height: 4}, want: 12},
		{name: "прямоугольник с дробными сторонами", shape: Rectangle{Width: 1.5, Height: 2}, want: 3},
		{name: "квадрат", shape: Square{Side: 5}, want: 25},
		{name: "квадрат с дробной стороной", shape: Square{Side: 0.5}, want: 0.25},
		{name: "круг нулевого радиуса", shape: Circle{Radius: 0}, want: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := tc.shape.Area(); math.Abs(got-tc.want) > eps {
				t.Errorf("%+v.Area() = %v, ожидалось %v", tc.shape, got, tc.want)
			}
		})
	}
}

// Все три типа обязаны удовлетворять Shape по значению.
func TestShapesSatisfyInterface(t *testing.T) {
	var _ Shape = Circle{}
	var _ Shape = Rectangle{}
	var _ Shape = Square{}
}

func TestTotalArea(t *testing.T) {
	testCases := []struct {
		name   string
		shapes []Shape
		want   float64
	}{
		{
			name:   "смешанный список",
			shapes: []Shape{Circle{Radius: 1}, Rectangle{Width: 2, Height: 3}, Square{Side: 1}},
			want:   math.Pi + 6 + 1,
		},
		{
			name:   "несколько фигур одного типа",
			shapes: []Shape{Square{Side: 2}, Square{Side: 3}},
			want:   13,
		},
		{
			name:   "одна фигура",
			shapes: []Shape{Rectangle{Width: 10, Height: 0.5}},
			want:   5,
		},
		{name: "nil-список", shapes: nil, want: 0},
		{name: "пустой список", shapes: []Shape{}, want: 0},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got := TotalArea(tc.shapes); math.Abs(got-tc.want) > eps {
				t.Errorf("TotalArea(%v) = %v, ожидалось %v", tc.shapes, got, tc.want)
			}
		})
	}
}

func TestMaxShape(t *testing.T) {
	testCases := []struct {
		name   string
		shapes []Shape
		want   Shape
	}{
		{
			name:   "максимум в середине списка",
			shapes: []Shape{Square{Side: 1}, Circle{Radius: 10}, Rectangle{Width: 1, Height: 2}},
			want:   Circle{Radius: 10},
		},
		{
			name:   "при равных площадях возвращается первая",
			shapes: []Shape{Rectangle{Width: 2, Height: 2}, Square{Side: 2}},
			want:   Rectangle{Width: 2, Height: 2},
		},
		{
			name:   "единственная фигура",
			shapes: []Shape{Square{Side: 7}},
			want:   Square{Side: 7},
		},
		{name: "nil-список", shapes: nil, want: nil},
		{name: "пустой список", shapes: []Shape{}, want: nil},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := MaxShape(tc.shapes)
			if tc.want == nil {
				if got != nil {
					t.Errorf("MaxShape(%v) = %v, ожидался nil", tc.shapes, got)
				}
				return
			}
			if got != tc.want {
				t.Errorf("MaxShape(%v) = %v, ожидалось %v", tc.shapes, got, tc.want)
			}
		})
	}
}

// TotalArea не должна ничего знать о типах из types.go: пользовательский
// тип с методом Area() обязан работать точно так же.
type triangle struct {
	base, height float64
}

func (tr triangle) Area() float64 { return tr.base * tr.height / 2 }

func TestTotalAreaWithForeignShape(t *testing.T) {
	shapes := []Shape{triangle{base: 4, height: 3}, Square{Side: 2}}

	if got := TotalArea(shapes); math.Abs(got-10) > eps {
		t.Errorf("TotalArea с типом из другого пакета = %v, ожидалось 10", got)
	}
}
