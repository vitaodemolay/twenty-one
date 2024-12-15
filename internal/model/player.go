package model

import (
	"strconv"
	"strings"
)

/*
=========================
Hand Section
=========================
*/

// hand Type
type hand []*Card

// Create a new Hand
func newHand() hand {
	return make(hand, 0)
}

// Get Hand in String format (ex: "[A♥, K♠]")
func (h hand) String() string {
	var cards []string
	for _, card := range h {
		cards = append(cards, card.Symbol())
	}
	return "[" + strings.Join(cards, ", ") + "]"
}

// Add a card to the hand
func (h *hand) AddCard(card *Card) {
	*h = append(*h, card)
}

/*
=========================
Player Section
=========================
*/

// Player struct
type Player struct {
	name  string
	hand  hand
	score int
}

// Create a new Player
func NewPlayer(name string) *Player {
	return &Player{
		name:  name,
		hand:  newHand(),
		score: 0,
	}
}

// Add a card to the player's hand
func (p *Player) AddCard(card *Card) {
	p.hand.AddCard(card)
	p.updateScore(card)
}

// update the player's score
func (p *Player) updateScore(card *Card) {
	p.score += card.Value()
}

// Get the player's score
func (p *Player) Score() int {
	return p.score
}

// Get Hand's player in string (ex: "David's Hand: [7♦, Q♣], Score: 17")
func (p *Player) String() string {
	return p.name + "'s Hand: " + p.hand.String() + ", Score: " + strconv.Itoa(p.score)
}

// Player new Round
func (p *Player) NewRound() {
	p.hand = newHand()
	p.score = 0
}
