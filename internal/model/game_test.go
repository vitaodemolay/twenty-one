package model

import (
	"testing"

	"github.com/vitaodemolay/twenty-one/internal/symbol"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	// Arrange & Act
	game := NewGame()

	// Assert
	assert.NotNil(t, game)
	assert.NotNil(t, game.deck)
	assert.Len(t, game.players, 0)
	assert.NotNil(t, game.dealer)
	assert.Equal(t, "Dealer", game.dealer.name)
	assert.Len(t, game.rounds, 0)
	assert.Equal(t, 0, game.currentRound)
}

func TestCreatePlayer(t *testing.T) {
	// Arrange
	game := NewGame()
	playerName := "Alice"

	// Act
	game.CreatePlayer(playerName)

	// Assert
	assert.Len(t, game.players, 1)
	assert.Equal(t, playerName, game.players[1].name)
}

func TestStartNewRound(t *testing.T) {
	// Arrange
	game := NewGame()
	game.CreatePlayer("Alice")

	// Act
	success, err := game.StartNewRound()

	// Assert
	assert.True(t, success)
	assert.Nil(t, err)
	assert.Equal(t, 1, game.currentRound)
	assert.Len(t, game.rounds, 1)
	assert.False(t, game.rounds[1].finished)
	assert.Equal(t, 52, game.deck.Len())
}

func TestStartNewRoundWithoutPlayers(t *testing.T) {
	// Arrange
	game := NewGame()

	// Act
	success, err := game.StartNewRound()

	// Assert
	assert.False(t, success)
	assert.EqualError(t, err, "not enough players to start a new round")
}

func TestStartNewRoundWithUnfinishedRound(t *testing.T) {
	// Arrange
	game := NewGame()
	game.CreatePlayer("Alice")
	game.StartNewRound()

	// Act
	success, err := game.StartNewRound()

	// Assert
	assert.False(t, success)
	assert.EqualError(t, err, "previous round has not ended yet")
}

func TestDealCardToPlayer(t *testing.T) {
	// Arrange
	game := NewGame()
	game.CreatePlayer("Alice")
	game.StartNewRound()
	player := game.players[1]

	// Act
	card, err := game.DealCardToPlayer(player)

	// Assert
	assert.NotNil(t, card)
	assert.Nil(t, err)
	assert.Len(t, player.hand, 1)
	assert.Equal(t, card, player.hand[0])
	assert.Equal(t, 51, game.deck.Len())
}

func TestDealCardToPlayerWithoutActiveRound(t *testing.T) {
	// Arrange
	game := NewGame()
	game.CreatePlayer("Alice")
	player := game.players[1]

	// Act
	card, err := game.DealCardToPlayer(player)

	// Assert
	assert.Nil(t, card)
	assert.EqualError(t, err, "no active round")
}

func TestCheckWhoIsWinner(t *testing.T) {
	// Arrange
	game := NewGame()
	game.CreatePlayer("Alice")
	game.CreatePlayer("Bob")
	game.StartNewRound()

	// Simulate a game state
	game.players[1].AddCard(NewCard(symbol.Hearts, symbol.King))
	game.players[1].AddCard(NewCard(symbol.Spades, symbol.Queen))
	game.players[2].AddCard(NewCard(symbol.Diamonds, symbol.Nine))
	game.players[2].AddCard(NewCard(symbol.Clubs, symbol.Eight))
	game.dealer.AddCard(NewCard(symbol.Hearts, symbol.Ten))
	game.dealer.AddCard(NewCard(symbol.Spades, symbol.Nine))

	// Act
	winner, err := game.CheckWhoIsWinner()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "Alice", winner.name)
	assert.Equal(t, 20, winner.Score())
}

func TestCheckWhoIsWinnerWithoutActiveRound(t *testing.T) {
	// Arrange
	game := NewGame()
	game.CreatePlayer("Alice")
	game.CreatePlayer("Bob")

	// Act
	winner, err := game.CheckWhoIsWinner()

	// Assert
	assert.Equal(t, Player{}, winner)
	assert.EqualError(t, err, "no active round")
}

func TestCheckWhoIsWinnerWithDealerWinBecauseHasDraw(t *testing.T) {
	// Arrange
	game := NewGame()
	game.CreatePlayer("Alice")
	game.CreatePlayer("Bob")
	game.StartNewRound()

	// Simulate a game state
	game.players[1].AddCard(NewCard(symbol.Hearts, symbol.King))
	game.players[1].AddCard(NewCard(symbol.Spades, symbol.Queen))
	game.players[2].AddCard(NewCard(symbol.Diamonds, symbol.King))
	game.players[2].AddCard(NewCard(symbol.Clubs, symbol.Queen))
	game.dealer.AddCard(NewCard(symbol.Hearts, symbol.Ten))
	game.dealer.AddCard(NewCard(symbol.Spades, symbol.Nine))

	// Act
	winner, err := game.CheckWhoIsWinner()

	// Assert
	assert.Nil(t, err)
	assert.Equal(t, "Dealer", winner.name)
	assert.Equal(t, 19, winner.Score())
}

func TestCheckWhoIsWinnerWithNoWinner(t *testing.T) {
	// Arrange
	game := NewGame()
	game.CreatePlayer("Alice")
	game.CreatePlayer("Bob")
	game.StartNewRound()

	// Simulate a game state
	game.players[1].AddCard(NewCard(symbol.Hearts, symbol.King))
	game.players[1].AddCard(NewCard(symbol.Spades, symbol.Queen))
	game.players[1].AddCard(NewCard(symbol.Spades, symbol.Queen))
	game.players[2].AddCard(NewCard(symbol.Diamonds, symbol.King))
	game.players[2].AddCard(NewCard(symbol.Clubs, symbol.Queen))
	game.players[2].AddCard(NewCard(symbol.Clubs, symbol.Queen))
	game.dealer.AddCard(NewCard(symbol.Hearts, symbol.Ten))
	game.dealer.AddCard(NewCard(symbol.Spades, symbol.Nine))
	game.dealer.AddCard(NewCard(symbol.Spades, symbol.Nine))

	// Act
	winner, err := game.CheckWhoIsWinner()

	// Assert
	assert.Equal(t, Player{}, winner)
	assert.EqualError(t, err, "no winner")
}

func TestCloseRoundWithWinner(t *testing.T) {
	// Arrange
	game := NewGame()
	game.CreatePlayer("Alice")
	game.StartNewRound()
	winner := game.players[1]

	// Act
	game.CloseRoundWithWinner(winner)

	// Assert
	assert.True(t, game.rounds[game.currentRound].finished)
	assert.Equal(t, *winner, game.rounds[game.currentRound].winner)
}
