package model

import (
	"testing"

	"github.com/vitaodemolay/twenty-one/internal/symbol"

	"github.com/stretchr/testify/assert"
)

func TestNewHand(t *testing.T) {
	// Arrange & Act
	hand := newHand()

	// Assert
	assert.NotNil(t, hand)
	assert.Len(t, hand, 0)
}

func TestHandString(t *testing.T) {
	// Arrange
	hand := newHand()
	hand.AddCard(NewCard(symbol.Hearts, symbol.Ace))
	hand.AddCard(NewCard(symbol.Spades, symbol.King))

	// Act
	result := hand.String()

	// Assert
	assert.Equal(t, "[A♥, K♠]", result)
}

func TestHandAddCard(t *testing.T) {
	// Arrange
	hand := newHand()
	card := NewCard(symbol.Diamonds, symbol.Queen)

	// Act
	hand.AddCard(card)

	// Assert
	assert.Len(t, hand, 1)
	assert.Contains(t, hand, card)
}

func TestNewPlayer(t *testing.T) {
	// Arrange
	playerName := "John Doe"

	// Act
	player := NewPlayer(playerName)

	// Assert
	assert.NotNil(t, player)
	assert.Equal(t, playerName, player.name)
	assert.Len(t, player.hand, 0)
	assert.Equal(t, 0, player.score)
}

func TestPlayerAddCard(t *testing.T) {
	// Arrange
	player := NewPlayer("Alice")
	card := NewCard(symbol.Hearts, symbol.Ten)

	// Act
	player.AddCard(card)

	// Assert
	assert.Len(t, player.hand, 1)
	assert.Contains(t, player.hand, card)
	assert.Equal(t, card.Value(), player.score)
}

func TestPlayerUpdateScore(t *testing.T) {
	// Arrange
	player := NewPlayer("Bob")
	card1 := NewCard(symbol.Clubs, symbol.Five)
	card2 := NewCard(symbol.Diamonds, symbol.King)

	// Act
	player.AddCard(card1)
	player.AddCard(card2)

	// Assert
	assert.Equal(t, 15, player.score)
}

func TestPlayerScore(t *testing.T) {
	// Arrange
	player := NewPlayer("Charlie")
	player.AddCard(NewCard(symbol.Spades, symbol.Ace))
	player.AddCard(NewCard(symbol.Hearts, symbol.Jack))

	// Act
	score := player.Score()

	// Assert
	assert.Equal(t, 11, score)
}

func TestPlayerString(t *testing.T) {
	// Arrange
	player := NewPlayer("David")
	player.AddCard(NewCard(symbol.Diamonds, symbol.Seven))
	player.AddCard(NewCard(symbol.Clubs, symbol.Queen))

	// Act
	result := player.String()

	// Assert
	assert.Equal(t, "David's Hand: [7♦, Q♣], Score: 17", result)
}

func TestPlayerNewRound(t *testing.T) {
	// Arrange
	player := NewPlayer("Emily")
	card1 := NewCard(symbol.Spades, symbol.Three)
	card2 := NewCard(symbol.Hearts, symbol.Nine)

	// Act
	player.AddCard(card1)
	player.AddCard(card2)
	player.NewRound()

	// Assert
	assert.Len(t, player.hand, 0)
	assert.Equal(t, 0, player.score)
}
