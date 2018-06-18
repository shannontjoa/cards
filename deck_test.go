package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Clubs" {
		t.Errorf("Expected first card of Ace of Clubs, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Spades" {
		t.Errorf("Expected last card of King of Spades, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	newDeck := newDeckFromFile("_decktesting")

	if len(newDeck) != 52 {
		t.Errorf("Expected number of cards of 52, but got %v", len(newDeck))
	}

	os.Remove("_decktesting")
}
