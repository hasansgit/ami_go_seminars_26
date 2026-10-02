package tickets

// Len возвращает число заявок в списке.
func (t ByUrgency) Len() int {
	panic("TODO: реализуйте ByUrgency.Len")
}

// Less сообщает, должна ли заявка с индексом i стоять раньше заявки
// с индексом j при сортировке по срочности (см. ByUrgency в types.go).
func (t ByUrgency) Less(i, j int) bool {
	panic("TODO: реализуйте ByUrgency.Less")
}

// Swap меняет местами заявки с индексами i и j.
func (t ByUrgency) Swap(i, j int) {
	panic("TODO: реализуйте ByUrgency.Swap")
}

// Sort сортирует заявки по срочности на месте, изменяя исходный слайс.
//
// После вызова заявки идут по возрастанию Priority, при равном приоритете —
// по возрастанию Arrived. Nil и пустой слайс допустимы, они не меняются.
func Sort(list []Ticket) {
	panic("TODO: реализуйте Sort")
}
