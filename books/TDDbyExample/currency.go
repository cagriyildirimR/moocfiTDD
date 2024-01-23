package TDDbyExample

import "log"

const (
	USDollar   = "USD"
	SwissFranc = "CHF"
)

type Wallet interface {
	Times(m int64) Wallet
	Equals(other Wallet) bool
	Amount() int64
	Type() string
	Add(m Wallet) Wallet
}

type Money struct {
	amount   int64
	currency string
}

func (a Money) Type() string {
	return a.currency
}

func (a Money) Amount() int64 {
	return a.amount
}

func (a Money) Times(x int64) Wallet {
	return Money{amount: a.Amount() * x, currency: a.currency}
}

func (a Money) Equals(other Wallet) bool {
	if a.Type() != other.Type() {
		exchangeRate, ok := exchangeRates[a.Type()+other.Type()]
		if !ok {
			log.Print("Exchange Rate doesn't exist")
		}
		return a.Amount() == int64(float32(other.Amount())*exchangeRate)
	}
	return a.Amount() == other.Amount()
}

func (a Money) Add(other Wallet) Wallet {
	if a.Type() == other.Type() {
		return Money{
			amount:   a.Amount() + other.Amount(),
			currency: a.Type(),
		}
	}

	exchangeRate, ok := exchangeRates[a.Type()+other.Type()]
	if !ok {
		log.Print("Exchange Rate doesn't exist")
	}

	return Money{amount: a.Amount() + int64(exchangeRate*float32(other.Amount())), currency: a.Type()}
}

var exchangeRates = make(map[string]float32)

func init() {
	exchangeRates[USDollar+SwissFranc] = 2
	exchangeRates[SwissFranc+USDollar] = 0.5
}
