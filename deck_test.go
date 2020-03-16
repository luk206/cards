package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 32 {
		t.Errorf("Expected deck length of 32 but got %v", len(d))
	}

	if d[0] != "One of Spades" {
		t.Errorf("Expected One of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "J of Clubs" {
		t.Errorf("Expected J of Clubs, but got %v", d[len(d)-1] )
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")

	deck := newDeck()
	deck.saveToFile("_deckTesting")
	loadedDeck := newDeckFromFile("_deckTesting")

	if len(loadedDeck) != 32 {
		t.Errorf("Expected deck length of 32 but got %v", len(loadedDeck))
	}

	os.Remove("_deckTesting")

}