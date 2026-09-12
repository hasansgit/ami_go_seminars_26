package greet

// Greet возвращает приветствие вида "Привет, Аня!".
// Если name — пустая строка, вернуть "Привет, мир!".
//
// Подсказка: строки склеиваются оператором +.
func Greet(name string) string {
	if len(name) == 0 {
		name = "мир"
	}
	return "Привет, " + name + "!"
}
