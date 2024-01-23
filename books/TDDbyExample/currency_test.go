package TDDbyExample

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoney_Equals(t *testing.T) {
	assert.True(t, Money{amount: 10, currency: USDollar}.Equals(Money{amount: 10, currency: USDollar}))
	assert.True(t, Money{amount: 10, currency: SwissFranc}.Equals(Money{amount: 10, currency: SwissFranc}))
	assert.False(t, Money{amount: 10, currency: USDollar}.Equals(Money{amount: 11, currency: SwissFranc}))
	assert.True(t, Money{amount: 10, currency: USDollar}.Equals(Money{amount: 5, currency: SwissFranc}))
	assert.False(t, Money{amount: 1, currency: TurkishLira}.Equals(Money{amount: 1, currency: USDollar}))
}

func TestMoney_Times(t *testing.T) {
	assert.True(t, Money{5, USDollar}.Times(2).Equals(Money{amount: 10, currency: USDollar}))
	assert.True(t, Money{amount: 5, currency: USDollar}.Times(3).Equals(Money{amount: 15, currency: USDollar}))
	assert.False(t, Money{amount: 5, currency: USDollar}.Times(2).Equals(Money{amount: 0, currency: USDollar}))
	assert.True(t, Money{amount: 5, currency: SwissFranc}.Times(2).Equals(Money{amount: 10, currency: SwissFranc}))
	assert.True(t, Money{amount: 5, currency: SwissFranc}.Times(3).Equals(Money{amount: 15, currency: SwissFranc}))
	assert.False(t, Money{amount: 5, currency: SwissFranc}.Times(2).Equals(Money{amount: 0, currency: SwissFranc}))
}

func TestMoney_Addition(t *testing.T) {
	AddDollars, err := Money{amount: 2, currency: USDollar}.Add(Money{amount: 1, currency: USDollar})
	assert.True(t, AddDollars.Equals(Money{amount: 3, currency: USDollar}))
	assert.Nil(t, err)
	AddFrancToDollars, err := Money{amount: 2, currency: USDollar}.Add(Money{amount: 1, currency: SwissFranc})
	assert.True(t, AddFrancToDollars.Equals(Money{amount: 4, currency: USDollar}))
	assert.Nil(t, err)
	AddDollarToFrancs, err := Money{amount: 2, currency: SwissFranc}.Add(Money{amount: 2, currency: USDollar})
	assert.True(t, AddDollarToFrancs.Equals(Money{amount: 3, currency: SwissFranc}))
	assert.Nil(t, err)

	AddLiraToDollar, err := Money{amount: 10, currency: USDollar}.Add(Money{amount: 10, currency: TurkishLira})
	assert.True(t, AddLiraToDollar.Equals(Money{}))
	assert.NotNil(t, err)
}
