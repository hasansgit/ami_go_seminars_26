// Этот файл менять не нужно — тесты рассчитаны именно на эти объявления.

package jsonreport

// Build — запись об одной сборке проекта.
type Build struct {
	Project string `json:"project"`
	Status  string `json:"status"`
	Seconds int    `json:"seconds"`
}

// Stat — сводная статистика по одному проекту.
type Stat struct {
	Project string `json:"project"`
	Total   int    `json:"total"`
	Failed  int    `json:"failed"`
	Seconds int    `json:"seconds"`
}
