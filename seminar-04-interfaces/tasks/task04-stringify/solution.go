package stringify

import (
	"fmt"
	"strconv"
	"strings"
)

// Stringify возвращает строковое представление произвольного значения.
//
// Правила представления по типу значения:
//   - nil — "<nil>";
//   - значение, реализующее fmt.Stringer, — результат его метода String();
//   - string — сама строка без изменений;
//   - bool — "true" или "false";
//   - int — десятичная запись, например 42 — это "42";
//   - float64 — ровно два знака после точки, например 1.5 — это "1.50";
//   - []string — элементы через запятую и пробел в квадратных скобках,
//     например []string{"a", "b"} — это "[a, b]", а пустой слайс — "[]";
//   - значение любого другого типа — "<unknown>".
func Stringify(v any) string {
	switch x := v.(type) {
	case nil:
		return "<nil>"
	case fmt.Stringer:
		return x.String()
	case string:
		return x
	case bool:
		return strconv.FormatBool(x)
	case int:
		return strconv.Itoa(x)
	case float64:
		return fmt.Sprintf("%.2f", x)
	case []string:
		return "[" + strings.Join(x, ", ") + "]"
	default:
		return "<unknown>"
	}
}
