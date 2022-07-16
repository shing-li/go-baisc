package main

import "fmt"

//custome type declaration

//creat a new type of "deck"
//which is a silce of strings
type deck []string

//建立第一個所以不需要receiver
func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}
	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}
	return cards
}

//reveiver
func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

/*
d:call這個function的實體
deck:所有type為deck的變數可以call this function
*/

func deal(d deck, handsize int) (deck, deck) { //回傳多個值
	return d[:handsize], d[handsize:]
}
