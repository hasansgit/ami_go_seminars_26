package tickets

import "sort"

// Len возвращает число заявок в списке.
func (t ByUrgency) Len() int {
	return len(t)
}

// Less сообщает, должна ли заявка с индексом i стоять раньше заявки
// с индексом j при сортировке по срочности (см. ByUrgency в types.go).
func (t ByUrgency) Less(i, j int) bool {
	return t[i].Priority < t[j].Priority || (t[i].Priority == t[j].Priority && t[i].Arrived < t[j].Arrived)
}

// Swap меняет местами заявки с индексами i и j.
func (t ByUrgency) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

// Sort сортирует заявки по срочности на месте, изменяя исходный слайс.
//
// После вызова заявки идут по возрастанию Priority, при равном приоритете —
// по возрастанию Arrived. Nil и пустой слайс допустимы, они не меняются.
func Sort(list []Ticket) {
	sort.Sort(ByUrgency(list))
}
