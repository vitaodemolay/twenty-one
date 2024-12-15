package model

import (
	"github.com/vitaodemolay/twenty-one/internal/symbol"
)

/*
=========================
Card Section
=========================
*/

// Card struct
type Card struct {
	nype  symbol.Nype
	ctype symbol.CardType
	value int
}

// Get the Card's Value (ex: A♠️ = 1, Q♠️ = 10)
func (c *Card) Value() int {
	return c.value
}

// Get the Card's String representation (ex: Ace of Spades)
func (c *Card) String() string {
	return c.ctype.String() + " of " + c.nype.String()
}

// Get the Card's Symbol representation (ex: A♠️)
func (c *Card) Symbol() string {
	return c.ctype.Symbol() + c.nype.Symbol()
}

// Function for Creating a new Card
func NewCard(n symbol.Nype, ct symbol.CardType) *Card {
	return &Card{nype: n, ctype: ct, value: cardValuesRule[ct]}
}
