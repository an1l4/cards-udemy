package main

func main() {
	cards := newDeck()
	cards.shuffle()
	cards.print()
	//cards := newDeckFromFile("my_cards")

	//cards := newDeck()
	//cards.saveToFile("my_cards")

	//fmt.Println(cards.toString())

	//cards.print()

	//hand, remainingCards := deal(cards, 5)

	//hand.print()
	//remainingCards.print()

}
