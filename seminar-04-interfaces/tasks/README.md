# Задачи семинара 4

Правите только `solution.go`. `solution_test.go` и `types.go` не трогаете:
тесты — это и есть условие, а типы зафиксированы, чтобы ваше решение
собиралось с чужим тестом. Прогнать задачи семинара — `make test-04`.
Подробно про цикл работы, команды и сдачу —
[`docs/how-to-solve.md`](../../docs/how-to-solve.md).

Условие каждой задачи целиком лежит в doc-комментариях над функциями в
`solution.go`. Читайте его внимательно: там оговорены граничные случаи, и
тесты проверяют каждый.

| # | Задача | Функции и методы | О чём |
|---|---|---|---|
| 1 | [`task01-money`](task01-money) | `Money.String` | Метод на своём типе; интерфейс `fmt.Stringer` |
| 2 | [`task02-shapes`](task02-shapes) | `Area` × 3, `TotalArea`, `MaxShape` | Неявное удовлетворение интерфейсу; функция над интерфейсом |
| 3 | [`task03-tickets`](task03-tickets) | `Len`, `Less`, `Swap`, `Sort` | Интерфейс `sort.Interface` целиком |
| 4 | [`task04-stringify`](task04-stringify) | `Stringify` | Пустой интерфейс `any`, type switch, проверка интерфейса |
| 5 | [`task05-censor`](task05-censor) | `New`, `Writer.Write` | Своя реализация `io.Writer`, композиция интерфейсов |

Пятая задача — дополнительная.

## Сигнатуры

```go
// task01-money
func (m Money) String() string

// task02-shapes
func (c Circle) Area() float64
func (r Rectangle) Area() float64
func (s Square) Area() float64
func TotalArea(shapes []Shape) float64
func MaxShape(shapes []Shape) Shape

// task03-tickets
func (t ByUrgency) Len() int
func (t ByUrgency) Less(i, j int) bool
func (t ByUrgency) Swap(i, j int)
func Sort(list []Ticket)

// task04-stringify
func Stringify(v any) string

// task05-censor
func New(dst io.Writer, banned []string) *Writer
func (w *Writer) Write(p []byte) (int, error)
```
