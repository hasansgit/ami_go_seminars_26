// error — это интерфейс. Свой тип ошибки и errors.As.
//
//	go run ./seminar-04-interfaces/examples/03-error-interface
package main

import (
	"errors"
	"fmt"
)

// Из пакета builtin:
//
//	type error interface {
//		Error() string
//	}
//
// Ошибка — это любое значение с методом Error() string. На семинаре 2
// были sentinel-ошибки (errors.New); они тоже просто реализуют этот
// интерфейс. Свой тип нужен, когда к ошибке хочется прикрепить данные.

// ValidationError — ошибка проверки поля: какое поле и что не так.
type ValidationError struct {
	Field  string
	Reason string
}

// Error делает ValidationError ошибкой. Получатель по указателю —
// стандарт для ошибок: их создают как &ValidationError{...}, и
// errors.As ниже рассчитан на *ValidationError.
func (e *ValidationError) Error() string {
	return fmt.Sprintf("поле %q: %s", e.Field, e.Reason)
}

func checkAge(age int) error {
	if age < 0 {
		return &ValidationError{Field: "age", Reason: "не может быть отрицательным"}
	}
	if age > 150 {
		return &ValidationError{Field: "age", Reason: "слишком большое"}
	}
	return nil
}

func main() {
	err := checkAge(-5)

	// errors.Is ищет совпадение по значению — это был семинар 2.
	// errors.As идёт дальше: ищет ошибку КОНКРЕТНОГО ТИПА и отдаёт её,
	// чтобы прочитать поля.
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Println("достали поля:", ve.Field, "-", ve.Reason)
	}

	// Ошибку, как и любой интерфейс, можно обернуть через %w —
	// errors.As найдёт ValidationError и под обёрткой.
	wrapped := fmt.Errorf("регистрация отклонена: %w", err)
	if errors.As(wrapped, &ve) {
		fmt.Println("под обёрткой тоже нашлось:", ve.Field)
	}

	// nil-ошибка — это интерфейс, равный nil. Проверка err != nil —
	// единственный каноничный способ.
	fmt.Println("корректный возраст:", checkAge(25) == nil) // true
}
