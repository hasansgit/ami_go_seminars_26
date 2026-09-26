# Задачи семинара 3

Правите только `solution.go`. `solution_test.go` и `types.go` не трогаете:
тесты — это и есть условие, а типы зафиксированы, чтобы ваше решение
собиралось с чужим тестом. Прогнать задачи семинара — `make test-03`.
Подробно про цикл работы, команды и сдачу —
[`docs/how-to-solve.md`](../../docs/how-to-solve.md).

Условие каждой задачи целиком лежит в doc-комментариях над функциями в
`solution.go`. Читайте его внимательно: там оговорены граничные случаи, и
тесты проверяют каждый.

| # | Задача | Функции | О чём |
|---|---|---|---|
| 1 | [`task01-segments`](task01-segments) | `Split`, `Cut` | Разрезать слайс на куски; вырезать из слайса диапазон |
| 2 | [`task02-setops`](task02-setops) | `Intersect`, `Difference` | Пересечение и разность двух списков строк |
| 3 | [`task03-fit`](task03-fit) | `Fit` | Обрезать строку под ограничение в байтах |
| 4 | [`task04-quota`](task04-quota) | `Apply` | Применить список операций расхода с проверкой лимитов |
| 5 | [`task05-jsonreport`](task05-jsonreport) | `Aggregate` | Свести JSON-массив записей о сборках по проектам |

Пятая задача — дополнительная.

## Сигнатуры

```go
// task01-segments
func Split(data []int, size int) ([][]int, error)
func Cut(data []int, from, to int) ([]int, error)

// task02-setops
func Intersect(a, b []string) []string
func Difference(a, b []string) []string

// task03-fit
func Fit(s string, limit int) (fitted string, truncated bool)

// task04-quota
func Apply(limits map[string]int, ops []Op) ([]Usage, error)

// task05-jsonreport
func Aggregate(data []byte) ([]byte, error)
```
