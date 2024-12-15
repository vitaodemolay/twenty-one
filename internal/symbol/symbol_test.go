package symbol

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNypeSymbol(t *testing.T) {
	tests := []struct {
		nype     Nype
		expected string
	}{
		{Hearts, "\u2665"},
		{Diamonds, "\u2666"},
		{Spades, "\u2660"},
		{Clubs, "\u2663"},
	}

	for _, test := range tests {
		// Arrange
		nype := test.nype

		// Act
		result := nype.Symbol()

		// Assert
		assert.Equal(t, test.expected, result, "Nype.Symbol() should return the correct symbol")
	}
}

func TestNypeString(t *testing.T) {
	tests := []struct {
		nype     Nype
		expected string
	}{
		{Hearts, "Hearts"},
		{Diamonds, "Diamonds"},
		{Spades, "Spades"},
		{Clubs, "Clubs"},
	}

	for _, test := range tests {
		// Arrange
		nype := test.nype

		// Act
		result := nype.String()

		// Assert
		assert.Equal(t, test.expected, result, "Nype.String() should return the correct string")
	}
}

func TestCardTypeSymbol(t *testing.T) {
	tests := []struct {
		cardType CardType
		expected string
	}{
		{Ace, "A"},
		{Two, "2"},
		{Three, "3"},
		{Four, "4"},
		{Five, "5"},
		{Six, "6"},
		{Seven, "7"},
		{Eight, "8"},
		{Nine, "9"},
		{Ten, "10"},
		{Queen, "Q"},
		{Jack, "J"},
		{King, "K"},
	}

	for _, test := range tests {
		// Arrange
		cardType := test.cardType

		// Act
		result := cardType.Symbol()

		// Assert
		assert.Equal(t, test.expected, result, "CardType.Symbol() should return the correct symbol")
	}
}

func TestCardTypeString(t *testing.T) {
	tests := []struct {
		cardType CardType
		expected string
	}{
		{Ace, "Ace"},
		{Two, "Two"},
		{Three, "Three"},
		{Four, "Four"},
		{Five, "Five"},
		{Six, "Six"},
		{Seven, "Seven"},
		{Eight, "Eight"},
		{Nine, "Nine"},
		{Ten, "Ten"},
		{Queen, "Queen"},
		{Jack, "Jack"},
		{King, "King"},
	}

	for _, test := range tests {
		// Arrange
		cardType := test.cardType

		// Act
		result := cardType.String()

		// Assert
		assert.Equal(t, test.expected, result, "CardType.String() should return the correct string")
	}
}
