package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Vocabulary struct {
	Word    string
	Meaning string
}

func main() {
	var vocabularyList []Vocabulary
	vocabularyList = append(vocabularyList, Vocabulary{Word: "fish", Meaning: "Fisch"})
	vocabularyList = append(vocabularyList, Vocabulary{Word: "tank", Meaning: "Panzer"})
	vocabularyList = append(vocabularyList, Vocabulary{Word: "flower", Meaning: "Blume"})
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Welcome to the Vocabulary Trainer!")
	for {
		fmt.Println("1. Add a new word")
		fmt.Println("2. Quiz me")
		fmt.Println("3. Exit")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter the word: ")
			scanner.Scan()
			word := scanner.Text()

			fmt.Print("Enter the meaning: ")
			scanner.Scan()
			meaning := scanner.Text()

			vocabularyList = append(vocabularyList, Vocabulary{Word: word, Meaning: meaning})
			fmt.Println("Word added successfully!")

		case "2":
			if len(vocabularyList) == 0 {
				fmt.Println("No words to quiz. Add some words first.")
				continue
			}

			fmt.Println("Starting the quiz...")
			correctCount := 0
			for _, vocab := range vocabularyList {
				fmt.Printf("What is the meaning of '%s'? ", vocab.Word)
				scanner.Scan()
				answer := scanner.Text()

				if strings.EqualFold(answer, vocab.Meaning) {
					fmt.Println("Correct!")
					correctCount++
				} else {
					fmt.Printf("Incorrect. The correct meaning is '%s'.\n", vocab.Meaning)
				}
			}
			fmt.Printf("You got %d out of %d correct.\n", correctCount, len(vocabularyList))

		case "3":
			fmt.Println("Exiting the Vocabulary Trainer. Goodbye!")
			return

		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
