package main

import (
	"testing"
	"os"
)


func TestNewDeck(t *testing.T){
	d := newDeck()

	if len(d)!= 52 {			//testing the length of the deck created
        t.Errorf("Expected 52 cards, got %d", len(d))		
    }
	if d[0] != "Ace of Spades"{		//testing the first card in the deck
		t.Errorf("Expected Ace of Spades, got %s", d[0])
	}
	if d[len(d)-1] != "King of Clubs"{		//testing the last card in the deck
		t.Errorf("Expected King of Clubs, got %s", d[len(d)-1])
	}
}
func TestSaveToDeckAndNewDeckFromFile (t *testing.T){ 		// testing if the save to file and deck from file functions are working
	os.Remove("_decktesting")
	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck:= newDeckFromFile("_decktesting")
	if len(loadedDeck)!= 52 {
        t.Errorf("Expected 52 cards, got %d", len(loadedDeck))
    }
	os.Remove("_decktesting")
}

