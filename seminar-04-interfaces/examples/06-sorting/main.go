// Сортировка: sort.Interface против slices.SortFunc.
//
//	go run ./seminar-04-interfaces/examples/06-sorting
package main

import (
	"fmt"
	"slices"
	"sort"
)

type Student struct {
	Name  string
	Score int
}

// Способ 1 (классический): объявить тип-слайс и реализовать
// sort.Interface — три метода Len, Less, Swap. Именно это вы будете
// делать в задаче 3.
type ByScoreDesc []Student

func (s ByScoreDesc) Len() int           { return len(s) }
func (s ByScoreDesc) Less(i, j int) bool { return s[i].Score > s[j].Score }
func (s ByScoreDesc) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func main() {
	group := []Student{
		{Name: "Аня", Score: 75},
		{Name: "Борис", Score: 92},
		{Name: "Вера", Score: 92},
		{Name: "Гриша", Score: 60},
	}

	// sort.Sort принимает интерфейс — ему безразличен конкретный тип.
	sort.Sort(ByScoreDesc(group))
	fmt.Println("по баллам (убывание):", group)

	// Способ 2 (современный, Go 1.21+): slices.SortFunc принимает
	// функцию сравнения — интерфейс не нужен вовсе. Функция возвращает
	// отрицательное число, если a раньше b; положительное — если позже;
	// 0 — если равны.
	slices.SortFunc(group, func(a, b Student) int {
		return a.Score - b.Score // по возрастанию баллов
	})
	fmt.Println("по баллам (возрастание):", group)

	// Когда что брать: сортировка одного типа по одному критерию
	// оформляется типом с тремя методами (её можно переиспользовать
	// и читать по имени), разовая сортировка — SortFunc.

	// sort.IsSorted проверяет отсортированность — удобно в тестах:
	fmt.Println("отсортировано по возрастанию:", slices.IsSortedFunc(group,
		func(a, b Student) int { return a.Score - b.Score }))
}
