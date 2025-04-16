package main

import "fmt"

// go is static type language like Java
// go is able to infer the type of a variable
func main() {
	// var card string = "Ace of Spades"
	// colon = is only used for variable declaration
	card := newCard()
	// card = "8 of Hearts"
	fmt.Println(card)	

}

func newCard() string {
	return "Five of Diamonds"
}