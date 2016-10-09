package main

import (
	"card/card"
	"fmt"
)

func main() {
	d := new(card.Deck)
	trials := 1000000
	c := 0
	for i := 0; i < trials; i++ {
		d.Init()
		for j := 0; j < 100; j++ {
			d.Shuffle()
		}
		//d.PrintDeck()
		c = c + countKings(d, i)
		printProgress(trials, i)
	}
	fmt.Printf("\r%.5f\n", float64(c)/float64(trials))
}

func countKings(d *card.Deck, round int) (count int) {
	count = 0
	for i := 0; i < 52; i++ {
		card := d.GetCard(i)
		if i < 51 {
			nextCard := d.GetCard(i + 1)
			if card.V == 13 && nextCard.V == 13 {
				//fmt.Printf("Round%d: %d/%s,%d/%s\n", round, card.V, card.S, nextCard.V, nextCard.S)
				count++
			}
		}
	}
	return
}

func printProgress(trials int, i int) {
	bucketSize := trials / 100
	if i%bucketSize == 0 {
		fmt.Printf("\r%d%%", (i/bucketSize + 1))
	}
}
