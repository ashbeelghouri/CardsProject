package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

// create a new type of deck which is a slice of string

type deck []string

var module = "deck"
var activeFunction = ""

var logger log = initialize(module, "none", []string{""})

func newDeck() deck {
	cards := deck{}
	logger = logger.setFunction("newDeck")
	logger.print([]string{"this is a simple function", "which can be used to log into system!"})
	cardSuits := []string{"spades", "diamonds", "hearts", "clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of  "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for i, card := range d {
		println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	logger = logger.setFunction("deal")
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	
	for i := range d {
		newPosition := r.Intn(len(d) - 1)
		d[i], d[newPosition] = d[newPosition], d[i]
	}
}
