package main

import "fmt"

func main() {
	cards := newDeckFromFile("my_cards")
	cards.print()
	fmt.Println("-----::After Shuffling::----")
	cards.shuffle()
	cards.print()
}
