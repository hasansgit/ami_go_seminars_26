package setops

import (
	"reflect"
	"testing"
)

func TestIntersect(t *testing.T) {
	testCases := []struct {
		name string
		a    []string
		b    []string
		want []string
	}{
		{
			name: "обычное пересечение",
			a:    []string{"a", "b", "c"},
			b:    []string{"b", "c", "d"},
			want: []string{"b", "c"},
		},
		{
			name: "повторы в a схлопываются",
			a:    []string{"x", "y", "x", "y"},
			b:    []string{"x"},
			want: []string{"x"},
		},
		{
			name: "повторы в b не влияют на результат",
			a:    []string{"x", "y"},
			b:    []string{"x", "x", "x"},
			want: []string{"x"},
		},
		{
			name: "порядок результата — порядок первых вхождений в a при другом порядке b",
			a:    []string{"c", "a", "b"},
			b:    []string{"a", "b", "c"},
			want: []string{"c", "a", "b"},
		},
		{
			name: "пустая строка — обычный элемент",
			a:    []string{"", "x"},
			b:    []string{""},
			want: []string{""},
		},
		{
			name: "a nil",
			a:    nil,
			b:    []string{"a"},
			want: nil,
		},
		{
			name: "a не nil, пустой",
			a:    []string{},
			b:    []string{"a"},
			want: []string{},
		},
		{
			name: "b nil означает пустое множество",
			a:    []string{"a", "b"},
			b:    nil,
			want: []string{},
		},
		{
			name: "результат пуст, но не nil",
			a:    []string{"a", "b"},
			b:    []string{"c", "d"},
			want: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Intersect(tc.a, tc.b)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Intersect(%#v, %#v) = %#v, ожидалось %#v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestDifference(t *testing.T) {
	testCases := []struct {
		name string
		a    []string
		b    []string
		want []string
	}{
		{
			name: "обычная разность",
			a:    []string{"a", "b", "c"},
			b:    []string{"b"},
			want: []string{"a", "c"},
		},
		{
			name: "повторы в a схлопываются",
			a:    []string{"x", "y", "x", "y"},
			b:    []string{"y"},
			want: []string{"x"},
		},
		{
			name: "повторы в b не влияют на результат",
			a:    []string{"x", "y"},
			b:    []string{"y", "y", "y"},
			want: []string{"x"},
		},
		{
			name: "порядок результата — порядок первых вхождений в a при другом порядке b",
			a:    []string{"c", "a", "b"},
			b:    []string{"z"},
			want: []string{"c", "a", "b"},
		},
		{
			name: "пустая строка — обычный элемент",
			a:    []string{"", "x"},
			b:    []string{"x"},
			want: []string{""},
		},
		{
			name: "a nil",
			a:    nil,
			b:    []string{"a"},
			want: nil,
		},
		{
			name: "a не nil, пустой",
			a:    []string{},
			b:    []string{"a"},
			want: []string{},
		},
		{
			name: "b nil означает пустое множество",
			a:    []string{"a", "b"},
			b:    nil,
			want: []string{"a", "b"},
		},
		{
			name: "результат пуст, но не nil",
			a:    []string{"a", "b"},
			b:    []string{"a", "b"},
			want: []string{},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got := Difference(tc.a, tc.b)
			if !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("Difference(%#v, %#v) = %#v, ожидалось %#v", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestIntersectDoesNotMutateInputs(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"b", "c", "d"}
	aCopy := append([]string(nil), a...)
	bCopy := append([]string(nil), b...)

	Intersect(a, b)

	if !reflect.DeepEqual(a, aCopy) {
		t.Fatalf("Intersect изменил a: получили %#v, ожидалось %#v", a, aCopy)
	}
	if !reflect.DeepEqual(b, bCopy) {
		t.Fatalf("Intersect изменил b: получили %#v, ожидалось %#v", b, bCopy)
	}
}

func TestDifferenceDoesNotMutateInputs(t *testing.T) {
	a := []string{"a", "b", "c"}
	b := []string{"b", "c", "d"}
	aCopy := append([]string(nil), a...)
	bCopy := append([]string(nil), b...)

	Difference(a, b)

	if !reflect.DeepEqual(a, aCopy) {
		t.Fatalf("Difference изменил a: получили %#v, ожидалось %#v", a, aCopy)
	}
	if !reflect.DeepEqual(b, bCopy) {
		t.Fatalf("Difference изменил b: получили %#v, ожидалось %#v", b, bCopy)
	}
}
