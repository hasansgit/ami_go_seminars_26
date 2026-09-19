package treetransform_test

import (
	"errors"
	"reflect"
	"slices"
	"strings"
	"testing"

	treetransform "github.com/DanielShinoda/ami_go_seminars_26/seminar-02-language/tasks/task05-treetransform"
)

func TestMapBuildsDeepCopyInPreOrder(t *testing.T) {
	source := sampleTree()
	before := sampleTree()
	var visited []int

	got, err := treetransform.Map(source, func(value int) (int, error) {
		visited = append(visited, value)
		return value * 10, nil
	})
	if err != nil {
		t.Fatalf("Map вернула ошибку: %v", err)
	}

	wantVisited := []int{1, 2, 3, 4}
	if !slices.Equal(visited, wantVisited) {
		t.Fatalf("порядок вызовов transform = %v, ожидалось %v", visited, wantVisited)
	}

	want := &treetransform.Node{
		Value: 10,
		Left:  &treetransform.Node{Value: 20},
		Right: &treetransform.Node{
			Value: 30,
			Left:  &treetransform.Node{Value: 40},
		},
	}
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("результат Map = %#v, ожидалось %#v", got, want)
	}

	assertNoSharedNodes(t, source, got)
	if !reflect.DeepEqual(source, before) {
		t.Fatalf("Map изменила исходное дерево: получено %#v, ожидалось %#v", source, before)
	}
}

func TestMapNilTree(t *testing.T) {
	called := false

	got, err := treetransform.Map(nil, func(value int) (int, error) {
		called = true
		return value, nil
	})
	if err != nil {
		t.Fatalf("Map(nil, transform) вернула ошибку: %v", err)
	}
	if got != nil {
		t.Fatalf("Map(nil, transform) = %#v, ожидалось nil", got)
	}
	if called {
		t.Fatal("transform был вызван для пустого дерева")
	}
}

func TestMapNilTransform(t *testing.T) {
	testCases := []struct {
		name   string
		root   *treetransform.Node
		before *treetransform.Node
	}{
		{name: "пустое дерево"},
		{name: "непустое дерево", root: sampleTree(), before: sampleTree()},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			got, err := treetransform.Map(tc.root, nil)
			if !errors.Is(err, treetransform.ErrNilTransform) {
				t.Fatalf("Map(root, nil) вернула ошибку %v, ожидалось ErrNilTransform", err)
			}
			if got != nil {
				t.Fatalf("Map(root, nil) = %#v, ожидалось nil", got)
			}
			if !reflect.DeepEqual(tc.root, tc.before) {
				t.Fatalf("Map(root, nil) изменила исходное дерево: получено %#v, ожидалось %#v", tc.root, tc.before)
			}
		})
	}
}

func TestMapWrapsErrorAndStops(t *testing.T) {
	source := sampleTree()
	before := sampleTree()
	errStop := errors.New("stop")
	var visited []int

	got, err := treetransform.Map(source, func(value int) (int, error) {
		visited = append(visited, value)
		if value == 3 {
			return 0, errStop
		}
		return value, nil
	})

	if got != nil {
		t.Fatalf("Map вернула частичное дерево %#v после ошибки", got)
	}
	if !errors.Is(err, errStop) {
		t.Fatalf("ошибка Map = %v, ожидалась обёрнутая errStop", err)
	}
	if err == nil || !strings.Contains(err.Error(), "3") {
		t.Fatalf("ошибка Map = %v, ожидался контекст со значением сбойного узла 3", err)
	}

	wantVisited := []int{1, 2, 3}
	if !slices.Equal(visited, wantVisited) {
		t.Fatalf("visited после ошибки = %v, ожидалось %v", visited, wantVisited)
	}
	if !reflect.DeepEqual(source, before) {
		t.Fatalf("Map изменила исходное дерево после ошибки: получено %#v, ожидалось %#v", source, before)
	}
}

func sampleTree() *treetransform.Node {
	return &treetransform.Node{
		Value: 1,
		Left:  &treetransform.Node{Value: 2},
		Right: &treetransform.Node{
			Value: 3,
			Left:  &treetransform.Node{Value: 4},
		},
	}
}

func assertNoSharedNodes(t *testing.T, source, mapped *treetransform.Node) {
	t.Helper()

	if source == nil || mapped == nil {
		if source != nil || mapped != nil {
			t.Fatalf("формы деревьев различаются: source=%#v mapped=%#v", source, mapped)
		}
		return
	}
	if source == mapped {
		t.Fatalf("результат переиспользует узел исходного дерева %p со значением %d", source, source.Value)
	}

	assertNoSharedNodes(t, source.Left, mapped.Left)
	assertNoSharedNodes(t, source.Right, mapped.Right)
}
