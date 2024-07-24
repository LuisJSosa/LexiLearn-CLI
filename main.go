package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"
)

// Word represents a vocabulary word and its definition
type Word struct {
	Term         string
	Definition   string
	CorrectCount int
	LastReview   time.Time
	NextReview   time.Time
}

var vocabList []Word
var fileName = "vocabLIst.json"

func main() {
	loadVocabList()
	defer saveVocabList()

	for {
		displayMenu()
		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 4.")
			os.Exit(1)
		}

		switch choice {
		case 1:
			fmt.Println("You chose to add a new vocabulary word.")
			addWord()
		case 2:
			fmt.Println("You chose to list all vocabulary words.")
			listWords()
		case 3:
			fmt.Println("You chose to review vocabulary words.")
			reviewWords()
		case 4:
			fmt.Println("Exiting... Goodbye!")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
			os.Exit(1)
		}
	}
}

func displayMenu() {
	fmt.Println("======================================")
	fmt.Println("         Welcome to LexiLearn         ")
	fmt.Println("======================================")
	fmt.Println("Please choose an option:")
	fmt.Println("1. Add a new vocabulary word")
	fmt.Println("2. List all vocabulary words")
	fmt.Println("3. Review vocabulary words")
	fmt.Println("4. Exit")
	fmt.Println("======================================")
	fmt.Print("Enter your choice: ")
}

// addWord prompts the user to enter a new word and its definition, and adds it to the vocabulary list
func addWord() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the term: ")
	term, _ := reader.ReadString('\n')
	term = strings.TrimSpace(term) // Remove trailing newline
	fmt.Print("Enter the definition: ")
	definition, _ := reader.ReadString('\n')
	definition = strings.TrimSpace(definition) // Remove trailing newline

	// Validate inputs
	if term == "" || definition == "" {
		fmt.Println("Both term and definition are required.")
		return
	}

	newWord := Word{
		Term:         term,
		Definition:   definition,
		CorrectCount: 0,
		LastReview:   time.Now(),
		NextReview:   time.Now().Add(24 * time.Hour),
	}
	vocabList = append(vocabList, newWord)

	fmt.Printf("Added: %s - %s\n", term, definition)
}

// listWords prints all the words in the vocabulary list
func listWords() {
	if len(vocabList) == 0 {
		fmt.Println("No vocabulary words found.")
		return
	}

	fmt.Println("Vocabulary List:")
	for _, word := range vocabList {
		fmt.Printf("%s: %s (Correct Count: %d)\n", word.Term, word.Definition, word.CorrectCount)
	}
}

// reviewWords allows the user to review vocabulary words
func reviewWords() {
	if len(vocabList) == 0 {
		fmt.Println("No vocabulary words found to review.")
		return
	}

	reader := bufio.NewReader(os.Stdin)
	for i := 0; i < len(vocabList); i++ {
		word := &vocabList[i]
		if time.Now().After(word.NextReview) {
			fmt.Printf("What is the definition of '%s'? ", word.Term)
			answer, _ := reader.ReadString('\n')
			answer = strings.TrimSpace(answer)

			if strings.EqualFold(answer, word.Definition) {
				word.CorrectCount++
				word.LastReview = time.Now()
				word.NextReview = calculateNextReview(word.CorrectCount)
				fmt.Println("Correct!")
			} else {
				word.CorrectCount = 0 // Reset correct count on failure
				word.LastReview = time.Now()
				word.NextReview = time.Now().Add(24 * time.Hour) // Review again tomorrow
				fmt.Printf("Incorrect. The correct definition is: %s\n", word.Definition)
			}
			fmt.Println("======================================")
		}
	}
}

func calculateNextReview(correctCount int) time.Time {
	switch correctCount {
	case 0:
		return time.Now().Add(24 * time.Hour) // 1 day later
	case 1:
		return time.Now().Add(3 * 24 * time.Hour) // 3 days later
	case 2:
		return time.Now().Add(7 * 24 * time.Hour) // 1 week later
	default:
		return time.Now().Add(30 * 24 * time.Hour) // 1 month later
	}
}

// saveVocabList saves the vocabulary list to a JSON file
func saveVocabList() {
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error saving vocabulary list:", err)
		return
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(vocabList)
	if err != nil {
		fmt.Println("Error encoding vocabulary list:", err)
	}
}

// loadVocabList loads the vocabulary list from a JSON file
func loadVocabList() {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return // No vocab list to load
		}
		fmt.Println("Error loading vocabulary list:", err)
		return
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&vocabList)
	if err != nil {
		fmt.Println("Error decoding vocabulary list:", err)
	}
}
