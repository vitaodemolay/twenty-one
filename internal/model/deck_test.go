package model

import (
	"testing"

	"github.com/vitaodemolay/twenty-one/internal/symbol"

	"github.com/stretchr/testify/assert"
)

func TestNewDeck(t *testing.T) {
	// Arrange
	expectedDeckSize := 52

	// Act
	deck := NewDeck()

	// Assert
	assert.NotNil(t, deck)
	assert.Len(t, deck.cards, expectedDeckSize)
	assert.Equal(t, expectedDeckSize, deck.Len())
}

func TestDeckIsEmpty(t *testing.T) {
	// Arrange
	deck := NewDeck()

	// Act & Assert
	assert.False(t, deck.IsEmpty())

	// Arrange
	emptyDeck := &Deck{}

	// Act & Assert
	assert.True(t, emptyDeck.IsEmpty())
}

func TestDeckLen(t *testing.T) {
	// Arrange
	deck := NewDeck()
	expectedLength := 52

	// Act
	length := deck.Len()

	// Assert
	assert.Equal(t, expectedLength, length)
}

func TestDeckShuffle(t *testing.T) {
	// Arrange
	deck := NewDeck()
	originalOrder := make([]*Card, len(deck.cards))
	copy(originalOrder, deck.cards)

	// Act
	deck.Shuffle()

	// Assert
	assert.NotEqual(t, originalOrder, deck.cards)
	assert.Len(t, deck.cards, len(originalOrder))
}

func TestDeckDeal(t *testing.T) {
	// Arrange
	deck := NewDeck()
	originalSize := deck.Len()

	// Act
	dealtCard := deck.Deal()

	// Assert
	assert.NotNil(t, dealtCard)
	assert.Equal(t, originalSize-1, deck.Len())
	assert.NotContains(t, deck.cards, dealtCard)
}

func TestDeckRenew(t *testing.T) {
	// Arrange
	deck := NewDeck()
	deck.Deal() // Remove one card

	// Act
	deck.Renew()

	// Assert
	assert.Len(t, deck.cards, 52)
	assert.Equal(t, 52, deck.Len())
}

func TestDeckContainsAllCards(t *testing.T) {
	// Arrange
	deck := NewDeck()

	// Act & Assert
	for n := symbol.Hearts; n <= symbol.Clubs; n++ {
		for ct := symbol.Ace; ct <= symbol.King; ct++ {
			expectedCard := NewCard(n, ct)
			assert.Contains(t, deck.cards, expectedCard)
		}
	}
}
