package main

import "fmt"

/*
golang basic types
1. bool
2. string
3. int
4. float64
*/

func main() {
	// var card string = "ace of spades"
	//:= only for declaration of variable not re-assigning value to it
	card := "Ace of spades"
	card = newCard()
	fmt.Println(card)

	//array = fixed length
	//slice = array that can grow or shrink and every element must be the same type aka homogeneous element

	//slice of type "string"
	// {} after the string = initialize value
	cards := []string{newCard(), "Ace of Diamonds"}
	fmt.Println(cards)
	cards = append(cards, "Seven of Diamonds") //appending the slice with a new value. It looks like ... operator in js
	//append also returns a new slice instead of modifying existing "cards" slice
	fmt.Println(cards)

	//i = index of current iteration
	//card = current value of current iteration
	// range cards = take the slice and iterates, keyword = "range" to iterate
	for i, card := range cards {
		fmt.Println(i, card)
	}
}

//this function return a type of string. return type is declared after () in function signature
func newCard() string {
	return "Five of Diamonds"
}
