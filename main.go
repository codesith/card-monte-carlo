package main

import (
	"card/card"
	"fmt"
)

func main() {
	d := new(card.Deck)
	trials := 10
	c := 0
	for i := 0; i < trials; i++ {
		d.Init()
		d.Shuffle()
		c = c + countKings(d)
	}
	fmt.Printf("Average times is %.2f\n", float64(c)/float64(trials))
}

func countKings(d *card.Deck) (count int) {
	count = 0
	for i := 0; i < 52; i++ {
		card := d.GetCard(i)
		if i < 51 {
			nextCard := d.GetCard(i + 1)
			if card.V == 13 && nextCard.V == 13 {
				fmt.Printf("%d%s and %d%s are next to each other\n", card.V, card.S, nextCard.V, nextCard.S)
				count++
			}
		}
	}
	return
}
