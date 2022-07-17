package main

//需要先go mod init <name>
import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d) != 16 {
		t.Errorf("not 16, is  %v", len(d))
	}
	if d[0] != "Ace of Spades" {
		t.Errorf("no, %v", d[0])
	}
	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("no, %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")
	if len(loadedDeck) != 16 {
		t.Errorf("not 16 is %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
