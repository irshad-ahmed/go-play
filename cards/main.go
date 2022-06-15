package main

func main() {
	//cards := deck{"5 of Spades", "Ace of Hearts"}
	//cards = append(cards, "6 of spades")
	//cards := newDeck()
	//	hand, remaingDeck := deal(cards, 5)
	//	cards.print()
	//	hand.print()
	//	remaingDeck.print()

	//fmt.Println(cards.toString())
	//cards.saveToFile("myCards.txt")
	cards := newDeckFromFile("myCards.txt")
	cards.shuffle()
	cards.print()
}

/* func main() {
	//var card string = "Ace of Spades"
	//card := "Ace of Spades"
	//card = "Five of diamonds"
	//	card := newCard()

	//cards := []string{"5 of Spades", newCard()}
	//cards = append(cards, "6 of spades")
	//fmt.Println(cards)

	/* 	for i, card := range cards {
		fmt.Println(i, card)
	}

	cards := deck{"5 of Spades", newCard()}
	cards = append(cards, "6 of spades")
	cards.print()
} */

/* func newCard() string {
	return "Five of Diamonds"
} */
