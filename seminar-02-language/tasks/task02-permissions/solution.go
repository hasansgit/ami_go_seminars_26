package permissions

// Grant возвращает права из current вместе со всеми правами из added.
//
// Функция не должна менять аргументы: current и added передаются по значению.
//
// Подсказка: чтобы включить биты, используйте оператор current | added.
func Grant(current, added Permission) Permission {
	panic("TODO: реализуйте Grant")
}

// Revoke удаляет из current все права из removed.
//
// Функция не должна менять аргументы: current и removed передаются по
// значению.
//
// Подсказка: чтобы сбросить биты, используйте оператор current &^ removed.
func Revoke(current, removed Permission) Permission {
	panic("TODO: реализуйте Revoke")
}

// Has возвращает true, только если в current присутствуют ВСЕ биты из
// required.
//
// Функция не должна менять аргументы: current и required передаются по
// значению.
//
// Подсказка: чтобы оставить только общие биты, используйте оператор
// current & required и сравните результат с required.
func Has(current, required Permission) bool {
	panic("TODO: реализуйте Has")
}
