package card

import (
	"fmt"
	"math/rand"
	"time"
)

type Deck struct {
	cards [52]Card
}

func (deck *Deck) Init() {
	deck.cards = [52]Card{}
	for i := 0; i < 13; i++ {
		deck.cards[i] = Card{i + 1, S}
		deck.cards[i+13] = Card{i + 1, H}
		deck.cards[i+26] = Card{i + 1, D}
		deck.cards[i+39] = Card{i + 1, C}
	}
}

var random = rand.New(rand.NewSource(time.Now().UnixNano()))

func (deck *Deck) Shuffle() {
	indices := random.Perm(52)
	for i := 0; i < 52; i++ {
		// swap
		c := deck.cards[i]
		deck.cards[i] = deck.cards[indices[i]]
		deck.cards[indices[i]] = c
	}
}

func (deck Deck) PrintDeck() {
	for i := 0; i < 52; i++ {
		fmt.Printf("%d%s, ", deck.cards[i].V, deck.cards[i].S)
	}
	fmt.Println()
}

func (deck Deck) GetCard(i int) (c Card) {
	return deck.cards[i]
}
