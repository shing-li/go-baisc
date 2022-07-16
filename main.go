package main

import "fmt"

func main() {
	//declare a slice
	//card := []string{"Three of spades", newCard()}
	//card := deck{}, 與上面相同，但需要同時執行deck.go
	cards := newDeck()

	hand, remainingCards := deal(cards, 5)
	hand.print()
	fmt.Println("=")
	remainingCards.print()
	//cards.print()

}

func newCard() string {
	return "Five of Diamonds"
}
