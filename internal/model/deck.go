package model

import (
	"github.com/vitaodemolay/twenty-one/internal/symbol"

	"golang.org/x/exp/rand"
)

/*
=========================
Deck Section
=========================
*/

// Deck struct
type Deck struct {
	cards []*Card
}

// Function to create a new deck
func NewDeck() *Deck {
	deck := &Deck{}
	for n := symbol.Hearts; n <= symbol.Clubs; n++ {
		for ct := symbol.Ace; ct <= symbol.King; ct++ {
			deck.cards = append(deck.cards, NewCard(n, ct))
		}
	}
	return deck
}

// Check if the deck is empty
func (d *Deck) IsEmpty() bool {
	return len(d.cards) == 0
}

// Get the number of cards in the deck
func (d *Deck) Len() int {
	return len(d.cards)
}

// Shuffle the deck
func (d *Deck) Shuffle() {
	for i := len(d.cards) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	}
}

// Deal a card from the deck
func (d *Deck) Deal() *Card {
	card := d.cards[len(d.cards)-1]
	d.cards = d.cards[:len(d.cards)-1]
	return card
}

// Renew the deck
func (d *Deck) Renew() {
	d.cards = nil
	d.cards = NewDeck().cards
}
