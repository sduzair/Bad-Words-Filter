package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type FlashCard struct {
	Term       string
	Definition string
}

func main() {
	// write your code here
	var term string
	var definition string

	scanner := bufio.NewScanner(os.Stdin)

	// get no of flashCards
	var noOfFlashcards int
	fmt.Println("Input the number of cards:")
	scanner.Scan()
	val, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	noOfFlashcards = int(val)

	// declare and allocate flashCards slice
	flashCards := make([]FlashCard, 0, noOfFlashcards)

	// get terms and definitions
	var flashCard FlashCard
	for i := 0; i < noOfFlashcards; i++ {
		fmt.Printf("The term for card #%d\n", i+1)
		scanner.Scan()
		term = scanner.Text()
		fmt.Printf("The definition for card #%d\n", i+1)
		scanner.Scan()
		definition = scanner.Text()
		flashCard.Term = term
		flashCard.Definition = definition

		flashCards = append(flashCards, flashCard)
	}

	for index := range flashCards {
		fmt.Printf("Print the definition of \"%s\"\n", flashCards[index].Term)
		scanner.Scan()
		if scanner.Text() == flashCards[index].Definition {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Wrong. The right answer is \"%s\"\n", flashCards[index].Definition)
		}
	}

}
