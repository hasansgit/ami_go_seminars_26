package segments

// Split делит data на последовательные куски длиной size; последний кусок
// может быть короче size, пустых кусков в результате нет.
//
// Каждый кусок делит общий backing array с data: изменение элемента куска
// видно через data, и наоборот. Ёмкость (cap) каждого куска равна его длине.
// Сам Split не меняет ни длину, ни содержимое data.
//
// Если size <= 0, результат — (nil, ErrInvalidSize).
// Для nil data результат — (nil, nil).
// Для не-nil пустого data результат — не-nil пустой [][]int и nil ошибка.
// Для непустого data число кусков в результате равно ceil(len(data)/size).
func Split(data []int, size int) ([][]int, error) {
	res := make([][]int, 0)
	if len(data) == 0 && data != nil {
		return res, nil
	}
	if size <= 0 {
		return nil, ErrInvalidSize
	}
	if data == nil {
		return nil, nil
	}

	ofs := 0

	for ofs+size < len(data) {
		sl := data[ofs : ofs+size : ofs+size]
		res = append(res, sl)
		ofs += size
	}

	sl := data[ofs:]
	res = append(res, sl)

	return res, nil
}

// Cut удаляет из data полуинтервал [from, to) и возвращает slice header
// длины len(data) - (to - from), в котором оставшиеся элементы сохраняют
// прежний порядок.
//
// Удаление происходит на месте: результат делит общий backing array с data.
// Элементы backing array в диапазоне от новой длины результата до прежней
// длины data после вызова равны нулю.
//
// Диапазон корректен при 0 <= from <= to <= len(data); в противном случае
// результат — (nil, ErrRange). При from == to содержимое не меняется, и
// возвращается slice header прежней длины. Для nil data корректен только
// случай from == to == 0, результат которого — (nil, nil). Для не-nil data
// результат не-nil, в том числе когда удалён весь слайс.
func Cut(data []int, from, to int) ([]int, error) {
	if data == nil {
		return nil, nil
	}
	if from < 0 || to < from || len(data) < to {
		return nil, ErrRange
	}

	copy(data[from:], data[to:])
	for i := len(data) - (to - from); i < len(data); i++ {
		data[i] = 0
	}
	return data[:len(data)-(to-from)], nil
}

// 0 1 2 3 4 5 6
// 0 1 2 4 5 6
