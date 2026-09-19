// Ошибки: sentinel-значения, обёртывание через %w и errors.Is.
//
// Запуск:
//
//	go run ./seminar-02-language/examples/07-errors
//
// Ошибка в Go — обычное значение, а не исключение. Её возвращают,
// а не бросают, и обрабатывают там же, где получили.
// Аналогия для тех, кто знает C++ или Python: это НЕ try/catch.
// Ближе всего — код возврата, только типизированный и с текстом внутри.
package main

import (
	"errors"
	"fmt"
)

// Sentinel-ошибки: заранее объявленные значения, по которым вызывающий код
// различает ситуации. Объявляются на уровне пакета, имя начинается с Err.
var (
	ErrNotFound   = errors.New("товар не найден")
	ErrOutOfStock = errors.New("товара нет на складе")
)

var stock = map[string]int{
	"кофе": 3,
	"чай":  0,
}

func main() {
	sentinels()
	wrapping()
	comparison()
}

func sentinels() {
	fmt.Println("--- sentinel-ошибки ---")

	for _, name := range []string{"кофе", "чай", "какао"} {
		count, err := take(name)
		switch {
		// errors.Is отвечает на вопрос «это та самая ошибка?».
		case errors.Is(err, ErrNotFound):
			fmt.Printf("  %-6s такого товара нет в каталоге\n", name)
		case errors.Is(err, ErrOutOfStock):
			fmt.Printf("  %-6s закончился\n", name)
		case err != nil:
			fmt.Printf("  %-6s неизвестная ошибка: %v\n", name, err)
		default:
			fmt.Printf("  %-6s выдано, осталось %d\n", name, count)
		}
	}
}

// take возвращает sentinel-ошибку как есть — без дополнительного контекста.
func take(name string) (int, error) {
	count, ok := stock[name]
	if !ok {
		return 0, ErrNotFound
	}
	if count == 0 {
		return 0, ErrOutOfStock
	}
	return count, nil
}

func wrapping() {
	fmt.Println("--- обёртывание через %w ---")

	err := processOrder("какао")

	// Текст накапливается слоями: каждый уровень добавил своё описание.
	fmt.Println("  текст ошибки:", err)

	// А errors.Is всё равно находит исходную ошибку в глубине цепочки.
	fmt.Println("  errors.Is(err, ErrNotFound):", errors.Is(err, ErrNotFound))

	// errors.Unwrap снимает ровно один слой — обычно им пользуются
	// не руками, а через errors.Is и errors.As.
	fmt.Println("  errors.Unwrap(err):", errors.Unwrap(err))
}

func processOrder(name string) error {
	if _, err := take(name); err != nil {
		// Глагол %w вкладывает исходную ошибку внутрь новой.
		// С %v текст был бы таким же, но errors.Is перестал бы работать —
		// это самая частая ошибка новичка.
		return fmt.Errorf("обработка заказа %q: %w", name, err)
	}
	return nil
}

func comparison() {
	fmt.Println("--- почему не ==  ---")

	wrapped := fmt.Errorf("слой: %w", ErrNotFound)

	// == сравнивает значения и видит только верхний слой.
	fmt.Println("  wrapped == ErrNotFound:       ", wrapped == ErrNotFound)
	// errors.Is разворачивает цепочку до конца.
	fmt.Println("  errors.Is(wrapped, ErrNotFound):", errors.Is(wrapped, ErrNotFound))

	// Отсюда правило: сравнивайте ошибки через errors.Is, а не через ==.
	// Иначе ваш код сломается в тот день, когда кто-то добавит контекст.
}
