package main

func main() {
	// var card string = "Ace of spades"
	cards := newDeck()
	cards.shuffle()
	cards.print()

	// hand, remainingCards := deal(cards, 5)

	// hand.print()
	// print("__")
	// remainingCards.print()

}
