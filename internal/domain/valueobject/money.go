package valueobject

import "fmt"

type Money struct {
	Amount   int
	Currency string
}

func NewMoney(amount int, currency string) Money {
	return Money{Amount: amount, Currency: currency}
}

func (m Money) String() string {
	return fmt.Sprintf("%d %s", m.Amount, m.Currency)
}
