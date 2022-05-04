package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type Flashcard struct {
	Term       string `json:"term"`
	Definition string `json:"definition"`
	Mistakes   int    `json:"mistakes"`
}

var flashCards []Flashcard

var logs []string

func main() {
	var filename string
	isExportFlashcardsWithFilename := false

	for _, arg := range os.Args[1:] {
		argKeyVal := strings.Split(arg, "=")
		if argKeyVal[0] == "--import_from" {
			filename = argKeyVal[1]
			importFlashcardsWithFilename(filename)
		} else if argKeyVal[0] == "--export_to" {
			filename = argKeyVal[1]
			isExportFlashcardsWithFilename = true
		} else {
			log.Panic("unknown cmd")
		}
	}

	scanner := bufio.NewScanner(os.Stdin)

loop:
	for {
		fmt.Println("Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):")
		logs = append(logs, "Input the action (add, remove, import, export, ask, exit, log, hardest card, reset stats):")
		scanner.Scan()
		switch scanner.Text() {
		case "add":
			addFlashcard(scanner)
		case "remove":
			removeFlashcard(scanner)
		case "import":
			importFlashcards(scanner)
		case "export":
			exportFlashcards(scanner)
		case "ask":
			test(scanner)
		case "exit":
			if isExportFlashcardsWithFilename {
				exportFlashcardsWithFilename(filename)
			}
			break loop
		case "log":
			writeLogs(scanner)
		case "hardest card":
			hardestCard(scanner)
		case "reset stats":
			resetStats(scanner)
		}
	}
	fmt.Println("bye bye")
}

func resetStats(scanner *bufio.Scanner) {
	for i := range flashCards {
		flashCards[i].Mistakes = 0
	}
	fmt.Printf("Card statistics have been reset.\n\n\n")
}

func hardestCard(scanner *bufio.Scanner) {
	maxMistakes := -1
	i := 0
	var hardestCardsIndex []int
	for i = range flashCards {
		if flashCards[i].Mistakes > maxMistakes {
			maxMistakes = flashCards[i].Mistakes
		}
	}
	if maxMistakes > 0 {
		for i := range flashCards {
			if flashCards[i].Mistakes == maxMistakes {
				hardestCardsIndex = append(hardestCardsIndex, i)
			}
		}
	}
	if len(hardestCardsIndex) == 1 {
		fmt.Printf("The hardest card is \"%s\". You have %d errors answering it.\n\n\n", flashCards[hardestCardsIndex[0]].Term, flashCards[hardestCardsIndex[0]].Mistakes)
	} else if len(hardestCardsIndex) > 1 {
		fmt.Printf("The hardest cards are \"%s\"", flashCards[hardestCardsIndex[0]].Term)
		for i := 1; i < len(hardestCardsIndex); i++ {
			fmt.Printf(",\"%s\" ", flashCards[hardestCardsIndex[i]].Term)
		}
		fmt.Print(".\n\n\n")
	} else {
		fmt.Print("There are no cards with errors.\n\n\n")
	}
}

func writeLogs(scanner *bufio.Scanner) {
	// open log file and insert logs
	fmt.Println("File name:")
	scanner.Scan()
	filename := scanner.Text()
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Panic(err)
	}
	defer file.Close()
	for i := range logs {
		fmt.Fprintln(file, logs[i])
	}
	fmt.Printf("The log has been saved.\n\n\n")
}

func test(scanner *bufio.Scanner) {
	logs = append(logs, "test")
	// get no of flashCards
	var noOfFlashcards int
	fmt.Println("How many times to ask?")
	logs = append(logs, "How many times to ask?")
	scanner.Scan()
	logs = append(logs, scanner.Text())
	val, _ := strconv.ParseInt(scanner.Text(), 10, 32)
	noOfFlashcards = int(val)

	for index := 0; index < noOfFlashcards; index++ {
		i := index % len(flashCards)
		fmt.Printf("Print the definition of \"%s\"\n", flashCards[i].Term)
		logs = append(logs, fmt.Sprintf("Print the definition of \"%s\"\n", flashCards[i].Term))
		scanner.Scan()
		userAns := scanner.Text()
		logs = append(logs, userAns)
		if userAns == flashCards[i].Definition {
			fmt.Printf("Correct!\n\n")
			logs = append(logs, "Correct!\n\n")
		} else if exists, index := isDefinitionExists(userAns); exists {
			fmt.Printf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\"\n\n", flashCards[i].Definition, flashCards[index].Term)
			logs = append(logs, fmt.Sprintf("Wrong. The right answer is \"%s\", but your definition is correct for \"%s\"\n\n", flashCards[i].Definition, flashCards[index].Term))
			flashCards[i].Mistakes++
		} else {
			fmt.Printf("Wrong. The right answer is \"%s\"\n\n", flashCards[i].Definition)
			logs = append(logs, fmt.Sprintf("Wrong. The right answer is \"%s\"\n\n", flashCards[i].Definition))
			flashCards[i].Mistakes++
		}
	}

}

func exportFlashcards(scanner *bufio.Scanner) {
	logs = append(logs, "export")
	fmt.Println("File name:")
	logs = append(logs, "File name:")
	scanner.Scan()
	filename := scanner.Text()
	logs = append(logs, filename)
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file.\n\n")
		logs = append(logs, "Error opening file.\n\n")
		return
	}
	defer file.Close()
	cardsSaved := 0
	for i, card := range flashCards {
		cardJson, err := json.Marshal(card)
		if err != nil {
			log.Panic(err)
		}
		_, wrError := fmt.Fprintln(file, string(cardJson))
		if wrError != nil {
			log.Panic(wrError)
		}
		cardsSaved = i + 1
	}
	fmt.Printf("%d cards have been saved.\n\n", cardsSaved)
	logs = append(logs, fmt.Sprintf("%d cards have been saved.\n\n", cardsSaved))

}

func exportFlashcardsWithFilename(filename string) {
	file, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("Error opening file.\n\n")
		return
	}
	defer file.Close()
	cardsSaved := 0
	for i, card := range flashCards {
		cardJson, err := json.Marshal(card)
		if err != nil {
			log.Panic(err)
		}
		_, wrError := fmt.Fprintln(file, string(cardJson))
		if wrError != nil {
			log.Panic(wrError)
		}
		cardsSaved = i + 1
	}
	fmt.Printf("%d cards have been saved.\n\n", cardsSaved)
}

func importFlashcards(scanner *bufio.Scanner) {
	logs = append(logs, "import")
	fmt.Println("File name:")
	logs = append(logs, "File name:")
	scanner.Scan()
	filename := scanner.Text()
	logs = append(logs, filename)
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("File not found.\n\n")
		logs = append(logs, "File not found.\n\n")

		return
	}
	defer file.Close()
	cardScanner := bufio.NewScanner(file)
	var card Flashcard
	cardsScanned := 0
	for cardScanner.Scan() {
		err := json.Unmarshal(cardScanner.Bytes(), &card)
		if err != nil {
			fmt.Println("error importing cards from file")
			logs = append(logs, "error importing cards from file")
			break
		}
		cardsScanned++
		if exists, index := isFlashcardExists(card); exists {
			// replace card in memory
			flashCards[index] = card
		} else {
			flashCards = append(flashCards, card)
		}
	}
	fmt.Printf("%d cards have been loaded.\n\n", cardsScanned)
	logs = append(logs, fmt.Sprintf("%d cards have been loaded.\n\n", cardsScanned))

}

func importFlashcardsWithFilename(filename string) {
	file, err := os.OpenFile(filename, os.O_RDONLY, 0644)
	if err != nil {
		fmt.Printf("File not found.\n\n")
		return
	}
	defer file.Close()
	cardScanner := bufio.NewScanner(file)
	var card Flashcard
	cardsScanned := 0
	for cardScanner.Scan() {
		err := json.Unmarshal(cardScanner.Bytes(), &card)
		if err != nil {
			fmt.Println("error importing cards from file")
			break
		}
		cardsScanned++
		if exists, index := isFlashcardExists(card); exists {
			// replace card in memory
			flashCards[index] = card
		} else {
			flashCards = append(flashCards, card)
		}
	}
	fmt.Printf("%d cards have been loaded.\n\n", cardsScanned)
}

func isFlashcardExists(card Flashcard) (bool, int) {
	for i, val := range flashCards {
		if val.Term == card.Term {
			return true, i
		}
	}
	return false, -1
}

func isCardExists(term string) (bool, int) {
	exists := false
	index := -1
	for i, card := range flashCards {
		if card.Term == term {
			exists = true
			index = i
		}
	}
	return exists, index
}

func removeFlashcard(scanner *bufio.Scanner) {
	logs = append(logs, "remove")
	fmt.Println("Which card?")
	logs = append(logs, "Which card?")
	scanner.Scan()
	term := scanner.Text()
	logs = append(logs, term)
	if val, i := isCardExists(term); val {
		// remove card
		flashCards[i] = flashCards[len(flashCards)-1]
		flashCards = flashCards[:len(flashCards)-1]
		fmt.Println("The card has been removed.")
		logs = append(logs, "The card has been removed.")
	} else {
		fmt.Printf("Can't remove \"%s\": there is no such card.\n\n", term)
		logs = append(logs, fmt.Sprintf("Can't remove \"%s\" there is no such card.\n\n", term))
	}

}

// func remove(s []any, i int) []any {
//     s[i] = s[len(s)-1]
//     return s[:len(s)-1]
// }

func isTermExists(term string) bool {
	exists := false
	for _, flashCard := range flashCards {
		if flashCard.Term == term {
			exists = true
		}
	}
	return exists
}

func isDefinitionExists(definition string) (bool, int) {
	for i, flashCard := range flashCards {
		if flashCard.Definition == definition {
			return true, i
		}
	}
	return false, -1
}

func addFlashcard(scanner *bufio.Scanner) {
	logs = append(logs, "add")
	var term string
	var definition string
	var flashCard Flashcard
	fmt.Println("The card:")
	logs = append(logs, "The card:")
	for {
		scanner.Scan()
		term = scanner.Text()
		logs = append(logs, term)
		if isTermExists(term) {
			fmt.Printf("The card \"%s\" already exists. Try again:\n", term)
			logs = append(logs, fmt.Sprintf("The card \"%s\" already exists. Try again:\n", term))
		} else {
			break
		}
	}
	fmt.Println("The definition of the card:")
	logs = append(logs, "The definition of the card:")
	for {
		scanner.Scan()
		definition = scanner.Text()
		logs = append(logs, definition)
		if exists, _ := isDefinitionExists(definition); exists {
			fmt.Printf("The definition \"%s\" already exists. Try again:\n", definition)
			logs = append(logs, fmt.Sprintf("The definition \"%s\" already exists. Try again:\n", definition))
		} else {
			break
		}
	}
	flashCard.Term = term
	flashCard.Definition = definition
	flashCards = append(flashCards, flashCard)
	fmt.Printf("The pair (\"%s\":\"%s\") has been added.\n\n", term, definition)
	logs = append(logs, fmt.Sprintf("The pair (\"%s\":\"%s\") has been added.\n\n", term, definition))

}

// func printFlashcards() {
// 	for _, val := range flashCards {
// 		fmt.Println(val)
// 	}
// }
