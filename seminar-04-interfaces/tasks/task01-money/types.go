// Этот файл менять не нужно — тесты рассчитаны именно на эти объявления.

package money

// Money — денежная сумма: Amount в минимальных единицах (копейках),
// Currency — код валюты, например "RUB".
type Money struct {
	Amount   int
	Currency string
}
