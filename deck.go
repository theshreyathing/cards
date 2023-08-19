package main

import (
	"fmt"
	"strings"
	"io/ioutil"
	"os"
	"math/rand"
	"time"
)

//create a new type of 'deck'
//which is a slice of string

type deck []string

func newDeck() deck { 	// collecting all the cards in the deck = 52
	cards:= deck {}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine","Ten","Jack", "Queen", "King"}

	for _, suit := range cardSuits{
		for _, value := range cardValues{
			cards = append(cards, value+ " of " +suit)
		}
	}
	return cards
}

func (d deck) print() { 					// Print function to print out all the cards in the slice one by one
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal (d deck, handSize int) (deck,deck){ 		// Dealing the card with the number of hand and splitting the deck into two
	return d[:handSize], d[handSize:]

}
func (d deck) toString() string{				// Converting the deck type to the string type ju joining all the elements and seperated by comma
	return strings.Join([]string(d), ",") 

}

func (d deck) saveToFile(filename string) error { 	// Converting the string to byte slice aand saving the file to hard drive
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)

}
func newDeckFromFile (filename string) deck{	// Reversing the above process and getting the saved file from local by converting it back to type deck
	bs, err:= ioutil.ReadFile(filename )
	if err != nil {		// error handling 
		// Option #1 = log the error and return a call to newDeck()
		// Option #2 = log the error and entirely quit the program
		fmt.Println("Error: ",err)
		os.Exit(1)
	}

	s := strings.Split(string(bs),",") 	// spliting the string to each card and converting it back to deck type
	return deck(s)
}
func (d deck) shuffle(){
	source := rand.NewSource(time.Now().UnixNano())
	r:= rand.New(source)
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i],d[newPosition] = d[newPosition], d[i]

	}
}