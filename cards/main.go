package main

import (
	"fmt"
)

func main() {
	// var card string = "Ace of Spades"
	// card:="Ace of Spades" //:= only to define a variabe
	// fmt.Println(card)
	var deckSize int
	deckSize = 52
	fmt.Println(deckSize)
	newCard:=newCard()
	fmt.Println(newCard)
	// slice()
	// decksImportReceiver()
	// newDeckFunWithReturn()
	// callFunctionWithReturnValues()
	// typeConversion()
	// fileSave()
	// fileRead()
	shuffleCards()
}

// return statements should define the data type
func newCard() string {
	return "Five of Diamonds"
}

func slice(){
	cards:=[]string{"Ace of Diamonds",newCard()} //type of slice string
	fmt.Println("Before appending",cards);
	//appending
	cards=append(cards,"Six of Spades") //creates a new one doesnt modify the old one
	fmt.Println("After appending",cards);

	//traversing through a slice
	for i,card:=range cards{
		fmt.Println(i,card)
	}
}
// Go is not a Object Oriented Language
func decksImportReceiver(){
	cards:=deck{"Ace of Diamonds",newCard()} //type of slice string
	fmt.Println("Before appending",cards);
	//appending
	cards=append(cards,"Six of Spades") //creates a new one doesnt modify the old one
	fmt.Println("After appending",cards);

	//traversing through a slice
	cards.print()
}

func newDeckFunWithReturn(){
	cards:=newDeck();
	cards.print()
}

func callFunctionWithReturnValues(){
	cards:=newDeck()
	hand,remainingCards:=deal(cards,5);
	hand.print()
	remainingCards.print()
}

func typeConversion(){
	greeting:="Hi there"
	fmt.Println([]byte(greeting))
	cards:=newDeck()
	fmt.Println(cards.toString())
}

func fileSave(){
	cards:=newDeck()
	cards.saveToFile("my_cards")
}

func fileRead(){
	cards:=newDeckFromFile("my_cards")
	cards.print()
}
func shuffleCards(){
	cards:=newDeck()
	cards.shuffle()
	cards.print()
}
