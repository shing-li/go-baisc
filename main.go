package main

import ()

func main() {
	//declare a slice
	//card := []string{"Three of spades", newCard()}
	//card := deck{}, 與上面相同，但需要同時執行deck.go
	// cards := newDeck()
	// fmt.Println(cards)
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("=")
	// remainingCards.print()
	//cards.print()
	// cards1 := newDeck()
	// a := cards1.toString()
	// ioutil.WriteFile("123.txt", []byte(a), 0644)
	// err := cards1.saveToFile("555.txt")
	// s, err := ioutil.ReadFile("555.txt")
	newDeck := newDeckFromFile("555.txt")

	newDeck.shuffle()
	newDeck.print()
}
