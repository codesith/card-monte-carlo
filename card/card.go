package card

type Suit string

const (
	S Suit = "Spades"
	H Suit = "Hearts"
	D Suit = "Dimonds"
	C Suit = "Clubs"
)

type Card struct {
	V int
	S Suit
}
