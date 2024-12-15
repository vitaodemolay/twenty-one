package model

import (
	"testing"

	"github.com/vitaodemolay/twenty-one/internal/symbol"

	"github.com/stretchr/testify/assert"
)

func TestNewCard(t *testing.T) {
	// Arrange
	nype := symbol.Hearts
	cardType := symbol.Ace

	// Act
	card := NewCard(nype, cardType)

	// Assert
	assert.NotNil(t, card, "NewCard() should not return nil")
	assert.Equal(t, nype, card.nype, "NewCard() should set the correct nype")
	assert.Equal(t, cardType, card.ctype, "NewCard() should set the correct card type")
	assert.Equal(t, 1, card.value, "NewCard() should set the correct value for Ace")
}

func TestCardValue(t *testing.T) {
	// Arrange
	card := NewCard(symbol.Spades, symbol.King)

	// Act
	value := card.Value()

	// Assert
	assert.Equal(t, 10, value, "Card.Value() should return the correct value for King")
}

func TestCardString(t *testing.T) {
	// Arrange
	card := NewCard(symbol.Diamonds, symbol.Queen)

	// Act
	str := card.String()

	// Assert
	assert.Equal(t, "Queen of Diamonds", str, "Card.String() should return the correct string representation")
}

func TestCardSymbol(t *testing.T) {
	// Arrange
	card := NewCard(symbol.Clubs, symbol.Jack)

	// Act
	symbol := card.Symbol()

	// Assert
	assert.Equal(t, "J♣", symbol, "Card.Symbol() should return the correct symbol representation")
}
