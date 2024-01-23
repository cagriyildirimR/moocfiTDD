package TDDbyExample

import "log"

const (
	USDollar   = "USD"
	SwissFranc = "CHF"
)

type Money struct {
	amount   int64
	currency string
}

func (a Money) Currency() string {
	return a.currency
}

func (a Money) Amount() int64 {
	return a.amount
}

func (a Money) Times(x int64) Money {
	return Money{amount: a.Amount() * x, currency: a.currency}
}

func (a Money) Equals(other Money) bool {
	if a.Currency() != other.Currency() {
		exchangeRate, ok := exchangeRates[a.Currency()+other.Currency()]
		if !ok {
			log.Print("Exchange Rate doesn't exist")
		}
		return a.Amount() == int64(float32(other.Amount())*exchangeRate)
	}
	return a.Amount() == other.Amount()
}

func (a Money) Add(other Money) Money {
	if a.Currency() == other.Currency() {
		return Money{
			amount:   a.Amount() + other.Amount(),
			currency: a.Currency(),
		}
	}

	exchangeRate, ok := exchangeRates[a.Currency()+other.Currency()]
	if !ok {
		log.Print("Exchange Rate doesn't exist")
	}

	return Money{amount: a.Amount() + int64(exchangeRate*float32(other.Amount())), currency: a.Currency()}
}

var exchangeRates = make(map[string]float32)

func init() {
	exchangeRates[USDollar+SwissFranc] = 2
	exchangeRates[SwissFranc+USDollar] = 0.5
}
