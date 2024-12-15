package model

import (
	"errors"

	"github.com/vitaodemolay/twenty-one/internal/symbol"
)

/*
=========================
Rules Section
=========================
*/

var cardValuesRule = map[symbol.CardType]int{
	symbol.Ace:   1,
	symbol.Two:   2,
	symbol.Three: 3,
	symbol.Four:  4,
	symbol.Five:  5,
	symbol.Six:   6,
	symbol.Seven: 7,
	symbol.Eight: 8,
	symbol.Nine:  9,
	symbol.Ten:   10,
	symbol.Queen: 10,
	symbol.Jack:  10,
	symbol.King:  10,
}

// Check player hand
func (p *Player) checkBrokenHisHand() bool {
	return p.Score() > 21
}

// Check if player draws
func (p *Player) hasDraw(g *Game) bool {
	if !p.isDealer() && p.Score() == g.dealer.Score() {
		return true
	}

	for _, player := range g.players {
		if p != player && p.Score() == player.Score() {
			return true
		}
	}

	return false
}

// Check if player is Dealer
func (p *Player) isDealer() bool {
	return p.name == "Dealer"
}

// Check if player has major score
func (p *Player) hasMajorScore(g *Game) bool {
	if !p.isDealer() && p.Score() <= g.dealer.Score() {
		return false
	}

	for _, player := range g.players {
		if p != player && p.Score() <= player.Score() {
			return false
		}
	}

	return true
}

/*
=========================
Round Section
=========================
*/

// Round struct
type Round struct {
	finished bool
	winner   Player
}

// Create a new Round
func NewRound() *Round {
	return &Round{
		finished: false,
	}
}

// Round add Winner
func (r *Round) AddWinner(winner *Player) {
	r.winner = *winner
	r.finished = true
}

/*
=========================
Game Section
=========================
*/

// Game struct
type Game struct {
	deck         *Deck
	players      map[int]*Player
	dealer       *Player
	rounds       map[int]*Round
	currentRound int
}

// Create a new Game
func NewGame() *Game {
	deck := NewDeck()

	players := make(map[int]*Player)
	rounds := make(map[int]*Round)

	return &Game{
		deck:         deck,
		players:      players,
		dealer:       NewPlayer("Dealer"),
		rounds:       rounds,
		currentRound: 0,
	}
}

func (g *Game) GetPlayerName(index int) (string, error) {
	if len(g.players) < 1 {
		return "", errors.New("not has players")
	}

	if index < 1 || index > len(g.players) {
		return "", errors.New("player not found")
	}

	return g.players[index].name, nil
}

// Game creates a new Player
func (g *Game) CreatePlayer(name string) {
	player := NewPlayer(name)
	g.players[len(g.players)+1] = player
}

// Game start a new Round
func (g *Game) StartNewRound() (bool, error) {
	if len(g.players) < 1 {
		return false, errors.New("not enough players to start a new round")
	}

	if g.currentRound > 0 && !g.rounds[g.currentRound].finished {
		return false, errors.New("previous round has not ended yet")
	}

	round := NewRound()
	g.currentRound += 1
	g.rounds[g.currentRound] = round

	g.dealer.NewRound()
	for _, player := range g.players {
		player.NewRound()
	}

	g.deck.Renew()
	g.deck.Shuffle()

	return true, nil
}

// Game Close Round with Winner
func (g *Game) CloseRoundWithWinner(winner *Player) {
	round := g.rounds[g.currentRound]
	round.AddWinner(winner)
}

// Game Deal Card for Player
func (g *Game) DealCardToPlayer(player *Player) (*Card, error) {
	if g.currentRound == 0 || g.rounds[g.currentRound].finished {
		return nil, errors.New("no active round")
	}

	card := g.deck.Deal()
	player.AddCard(card)

	if player.checkBrokenHisHand() {
		return card, errors.New("the player bursted out")
	}

	return card, nil
}

// Check Who is Winner
func (g *Game) CheckWhoIsWinner() (Player, error) {
	if g.currentRound == 0 || g.rounds[g.currentRound].finished {
		return Player{}, errors.New("no active round")
	}

	var winner *Player

	if !g.dealer.checkBrokenHisHand() && (g.dealer.hasMajorScore(g) || g.dealer.hasDraw(g)) {
		winner = g.dealer

	} else {
		for _, player := range g.players {
			if !g.dealer.checkBrokenHisHand() && player.hasDraw(g) {
				winner = g.dealer
				break
			}
			if !player.checkBrokenHisHand() && player.hasMajorScore(g) {
				winner = player
				break
			}
		}
	}

	if winner == nil {
		return Player{}, errors.New("no winner")
	}

	return *winner, nil
}
