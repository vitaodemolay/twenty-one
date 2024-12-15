package symbol

/*
=========================
Nype Section
=========================
*/

// Nype enum type
type Nype int

// Enums of Nypes
const (
	Hearts Nype = iota
	Diamonds
	Spades
	Clubs
)

// Get the Nype symbol
func (n Nype) Symbol() string {
	return [...]string{"\u2665", "\u2666", "\u2660", "\u2663"}[n]
}

// Get the Nype in string
func (n Nype) String() string {
	return [...]string{"Hearts", "Diamonds", "Spades", "Clubs"}[n]
}

/*
=========================
CardType Section
=========================
*/

// CardType enum type
type CardType int

// Enums of CardType
const (
	Ace CardType = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Queen
	Jack
	King
)

// Get the CardType symbol
func (c CardType) Symbol() string {
	return [...]string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Q", "J", "K"}[c]
}

// Get The CardType in string
func (c CardType) String() string {
	return [...]string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Queen", "Jack", "King"}[c]
}
