package TDDbyExample

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMoney_Times(t *testing.T) {
	assert.True(t, Currency{5, USDollar}.Times(2).Equals(Currency{amount: 10, name: USDollar}))
	assert.True(t, Currency{amount: 5, name: USDollar}.Times(3).Equals(Currency{amount: 15, name: USDollar}))
	assert.False(t, Currency{amount: 5, name: USDollar}.Times(2).Equals(Currency{amount: 0, name: USDollar}))
	assert.True(t, Currency{amount: 5, name: SwissFranc}.Times(2).Equals(Currency{amount: 10, name: SwissFranc}))
	assert.True(t, Currency{amount: 5, name: SwissFranc}.Times(3).Equals(Currency{amount: 15, name: SwissFranc}))
	assert.False(t, Currency{amount: 5, name: SwissFranc}.Times(2).Equals(Currency{amount: 0, name: SwissFranc}))
}

func TestMoney_Equals(t *testing.T) {
	assert.True(t, Currency{amount: 10, name: USDollar}.Equals(Currency{amount: 10, name: USDollar}))
	assert.True(t, Currency{amount: 10, name: SwissFranc}.Equals(Currency{amount: 10, name: SwissFranc}))
	assert.False(t, Currency{amount: 10, name: USDollar}.Equals(Currency{amount: 11, name: SwissFranc}))
	assert.True(t, Currency{amount: 10, name: USDollar}.Equals(Currency{amount: 5, name: SwissFranc}))
}

func TestMoney_Addition(t *testing.T) {
	assert.True(t, Currency{amount: 2, name: USDollar}.Add(Currency{amount: 1, name: USDollar}).Equals(Currency{amount: 3, name: USDollar}))
	assert.True(t, Currency{amount: 2, name: USDollar}.Add(Currency{amount: 1, name: SwissFranc}).Equals(Currency{amount: 4, name: USDollar}))
	assert.True(t, Currency{amount: 2, name: SwissFranc}.Add(Currency{amount: 2, name: USDollar}).Equals(Currency{amount: 3, name: SwissFranc}))
	assert.True(t, Currency{amount: 10, name: USDollar}.Add(Currency{amount: 10, name: SwissFranc}).Equals(Currency{amount: 30, name: USDollar}))
}
