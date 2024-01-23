package TDDbyExample

import "log"

const (
	USDollar   = "USD"
	SwissFranc = "CHF"
)

type Money interface {
	Times(m int64) Money
	Equals(other Money) bool
	Amount() int64
	Type() string
	Add(m Money) Money
}

type Currency struct {
	amount int64
	name   string
}

func (a Currency) Type() string {
	return a.name
}

func (a Currency) Amount() int64 {
	return a.amount
}

func (a Currency) Times(x int64) Money {
	return Currency{amount: a.Amount() * x, name: a.name}
}

func (a Currency) Equals(other Money) bool {
	if a.Type() != other.Type() {
		exchangeRate, ok := exchangeRates[a.Type()+other.Type()]
		if !ok {
			log.Print("Exchange Rate doesn't exist")
		}
		return a.Amount() == int64(float32(other.Amount())*exchangeRate)
	}
	return a.Amount() == other.Amount()
}

func (a Currency) Add(other Money) Money {
	if a.Type() == other.Type() {
		return Currency{
			amount: a.Amount() + other.Amount(),
			name:   a.Type(),
		}
	}

	exchangeRate, ok := exchangeRates[a.Type()+other.Type()]
	if !ok {
		log.Print("Exchange Rate doesn't exist")
	}

	return Currency{amount: a.Amount() + int64(exchangeRate*float32(other.Amount())), name: a.Type()}
}

var exchangeRates = make(map[string]float32)

func init() {
	exchangeRates[USDollar+SwissFranc] = 2
	exchangeRates[SwissFranc+USDollar] = 0.5
}
