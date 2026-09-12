package minmax

// MinMax возвращает минимальный и максимальный элементы слайса,
// а также признак ok = true, если слайс не пуст.
// Для пустого (или nil) слайса вернуть 0, 0, false.
//
// Это задача про МНОЖЕСТВЕННЫЙ ВОЗВРАТ — идиому Go.
// Возвращать признак успеха отдельным значением здесь нормально и правильно.
//
// Подсказка: перебрать элементы можно так:
//
//	for _, v := range nums { ... }
func MinMax(nums []int) (int, int, bool) {
	if len(nums) == 0 {
		return 0, 0, false
	}

	mn, mx := nums[0], nums[0]
	for _, v := range nums {
		if v > mx {
			mx = v
		}
		if v < mn {
			mn = v
		}
	}
	return mn, mx, true
}
