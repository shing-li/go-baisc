package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

//custome type declaration
//byte slice => string，用ascii code存string

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

//把deck轉回slice of string, []string()
//接著把slice of string轉換成一個string，之後再轉成[]byte以寫入檔案
//deck -> []string -> string -> []byte
func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	s := d.toString()
	return ioutil.WriteFile(filename, []byte(s), 0666) //filemode 0666 means everyone can Read and Write
}

// []byte -> string -> []string -> deck
func newDeckFromFile(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	if err != nil { //error handling in golang
		fmt.Println("error", err)
		os.Exit(1)
	}
	sliceOfString := strings.Split(string(byteSlice), ",") //[]byte -> string -> []string
	return deck(sliceOfString)                             //[]string -> deck
}

func (d deck) shuffle() {
	//needs a int64 source, use time.UnixNano to return the current time as seed
	source := rand.NewSource(time.Now().UnixNano())
	n := rand.New(source)
	for i := range d {
		// newPosition := rand.Intn(len(d) - 1)	沒有初始化seed
		newPosition := n.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i] //one line swap
	}
}
