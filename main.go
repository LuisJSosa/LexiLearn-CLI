package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"os"
	"sort"
	"strconv"
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
var deckName = "default"
var fileName = deckName + "vocabList.json"

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}

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
			addWord()
		case 2:
			listWords()
		case 3:
			reviewWords()
		case 4:
			manageDecks()
		case 5:
			fmt.Println("Exiting LexiLearn... Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 4.")
			time.Sleep(2 * time.Second)
		}
	}
}

func displayMenu() {
	fmt.Println("======================================")
	fmt.Println("         Welcome to LexiLearn         ")
	fmt.Println("======================================")
	fmt.Println("Current Deck:", deckName)
	fmt.Println("Please choose an option:")
	fmt.Println("1. Add a new vocabulary word")
	fmt.Println("2. List all vocabulary words")
	fmt.Println("3. Review vocabulary words")
	fmt.Println("4. Manage decks")
	fmt.Println("4. Exit")
	fmt.Println("======================================")
	fmt.Print("Enter your choice: ")
}

// addWord prompts the user to enter a new word and its definition, and adds it to the vocabulary list
func addWord() {
	reader := bufio.NewReader(os.Stdin)
	clearScreen()
	color.Cyan("You chose to add a new vocabulary word.")
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
		NextReview:   time.Now().Add(-24 * time.Hour),
	}
	vocabList = append(vocabList, newWord)

	color.Green("Added: %s - %s\n", term, definition)
	time.Sleep(2 * time.Second)
}

// listWords prints all the words in the vocabulary list
func listWords() {
	clearScreen()

	if len(vocabList) == 0 {
		color.Red("No vocabulary words found.")
		return
	}
	color.Blue("Vocabulary List:")
	for _, word := range vocabList {
		fmt.Printf("- %s: %s (Next review: %s)\n", word.Term, word.Definition, word.NextReview.Format("01-02-2006 03:04 PM"))
	}
	time.Sleep(5 * time.Second)
}

// reviewWords allows the user to review vocabulary words
func reviewWords() {
	if len(vocabList) == 0 {
		color.Red("No vocabulary words found to review.")
		time.Sleep(2 * time.Second)
		return
	}

	reader := bufio.NewReader(os.Stdin)
	startTime := time.Now()

	for {
		clearScreen()
		sort.Slice(vocabList, func(i, j int) bool {
			return vocabList[i].NextReview.Before(vocabList[j].NextReview)
		})

		reviewedAny := false

		for i := 0; i < len(vocabList); i++ {
			word := &vocabList[i]

			// Show the word even if it's not yet due for review
			color.Yellow("Think of the definition of '%s'. Press Enter to see the correct definition.", word.Term)
			_, _ = reader.ReadString('\n')

			color.Cyan("The correct definition is: %s\n", word.Definition)

			var confidence int
			for {
				fmt.Print("How confident do you feel about this word? (1-5): ")
				input, err := reader.ReadString('\n')
				if err == nil {
					input = strings.TrimSpace(input)
					confidence, err = strconv.Atoi(input)
					if err == nil && confidence >= 1 && confidence <= 5 {
						break
					}
				}
				color.Red("Invalid input. Please enter a number between 1 and 5.")
			}

			word.LastReview = time.Now()
			word.NextReview = calculateNextReview(confidence)
			color.Cyan("Next review for '%s' scheduled.\n", word.Term)
			color.Cyan("======================================")

			reviewedAny = true

			if time.Since(startTime) > 10*time.Minute {
				color.Red("Review session has been going on for 10 minutes. Taking a break!")
				time.Sleep(2 * time.Second)
				return
			}
		}

		// If no words need review, ask user if they want to review all words again
		if !reviewedAny {
			color.Green("No words need to be reviewed right now.")
			fmt.Print("Do you want to review all words again? (yes/no): ")
			cont, _ := reader.ReadString('\n')
			cont = strings.TrimSpace(cont)
			if strings.ToLower(cont) == "yes" {
				// Review all words again
				continue
			} else {
				time.Sleep(2 * time.Second)
				break
			}
		} else {
			// Ask user if they want to continue reviewing
			fmt.Print("Do you want to continue reviewing? (yes/no): ")
			cont, _ := reader.ReadString('\n')
			cont = strings.TrimSpace(cont)
			if strings.ToLower(cont) != "yes" {
				break
			}
		}
	}
}

// calculateNextReview determines the next review time based on confidence level
func calculateNextReview(confidence int) time.Time {
	switch confidence {
	case 0:
		return time.Now().Add(1 * time.Minute) // 1 minute later
	case 1:
		return time.Now().Add(3 * time.Minute) // 3 minutes later
	case 2:
		return time.Now().Add(5 * time.Minute) // 5 minutes later
	case 3:
		return time.Now().Add(10 * time.Minute) // 10 minutes later
	case 4:
		return time.Now().Add(15 * time.Minute) // 15 minutes later
	case 5:
		return time.Now().Add(25 * time.Minute) // 25 minutes later
	default:
		return time.Now().Add(2 * time.Hour) // 2 hours later
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
	} else {
		color.Green("Vocabulary list saved successfully!")
	}
}

// loadVocabList loads the vocabulary list from a JSON file
func loadVocabList() {
	file, err := os.Open(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			color.Yellow("No previous vocabulary list found. Starting fresh!")
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
	} else {
		color.Green("Vocabulary list loaded successfully!")
	}
}

func manageDecks() {
	for {
		clearScreen()
		fmt.Println("======================================")
		fmt.Println("           Manage Decks               ")
		fmt.Println("======================================")
		fmt.Println("1. List decks")
		fmt.Println("2. Create a new deck")
		fmt.Println("3. Select a deck")
		fmt.Println("4. Rename a deck")
		fmt.Println("5. Back to main menu")
		fmt.Println("======================================")
		fmt.Print("Enter your choice: ")

		var choice int
		_, err := fmt.Scanf("%d", &choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter a number between 1 and 5.")
			continue
		}

		switch choice {
		case 1:
			listDecks()
		case 2:
			createDeck()
		case 3:
			selectDeck()
		case 4:
			renameDeck()
		case 5:
			return
		default:
			fmt.Println("Invalid choice. Please enter a number between 1 and 5.")
			time.Sleep(2 * time.Second)
		}
	}
}

func listDecks() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading decks:", err)
		return
	}

	color.Blue("Available Decks:")
	for i, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			fmt.Printf("%d. %s\n", i+1, strings.TrimSuffix(file.Name(), ".json"))
		}
	}
	time.Sleep(5 * time.Second)
}

func createDeck() {
	reader := bufio.NewReader(os.Stdin)
	clearScreen()
	color.Cyan("You chose to create a new deck.")
	fmt.Print("Enter the name of the new deck: ")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	if name == "" {
		color.Red("Deck name cannot be empty.")
		time.Sleep(2 * time.Second)
		return
	}

	fileName = name + ".json"
	vocabList = []Word{}
	saveVocabList()
	color.Green("Deck '%s' created successfully!", name)
	time.Sleep(2 * time.Second)
}

func selectDeck() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading decks:", err)
		return
	}

	color.Blue("Available Decks:")
	deckFiles := []string{}
	for i, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			deckFiles = append(deckFiles, file.Name())
			fmt.Printf("%d. %s\n", i+1, strings.TrimSuffix(file.Name(), ".json"))
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of the deck to select: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(deckFiles) {
		color.Red("Invalid choice.")
		time.Sleep(2 * time.Second)
		return
	}

	fileName = deckFiles[choice-1]
	deckName = strings.TrimSuffix(fileName, ".json")
	loadVocabList()
	color.Green("Deck '%s' selected successfully!", deckName)
	time.Sleep(2 * time.Second)
}

func renameDeck() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error reading decks:", err)
		return
	}

	color.Blue("Available Decks:")
	deckFiles := []string{}
	for i, file := range files {
		if strings.HasSuffix(file.Name(), ".json") {
			deckFiles = append(deckFiles, file.Name())
			fmt.Printf("%d. %s\n", i+1, strings.TrimSuffix(file.Name(), ".json"))
		}
	}

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the number of the deck to rename: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	choice, err := strconv.Atoi(input)
	if err != nil || choice < 1 || choice > len(deckFiles) {
		color.Red("Invalid choice.")
		time.Sleep(2 * time.Second)
		return
	}

	oldFileName := deckFiles[choice-1]
	oldDeckName := strings.TrimSuffix(oldFileName, ".json")

	fmt.Print("Enter the new name for the deck: ")
	newName, _ := reader.ReadString('\n')
	newName = strings.TrimSpace(newName)
	if newName == "" {
		color.Red("Deck name cannot be empty.")
		time.Sleep(2 * time.Second)
		return
	}

	newFileName := newName + ".json"
	err = os.Rename(oldFileName, newFileName)
	if err != nil {
		fmt.Println("Error renaming deck:", err)
		return
	}

	if deckName == oldDeckName {
		deckName = newName
		fileName = newFileName
	}

	color.Green("Deck '%s' renamed to '%s' successfully!", oldDeckName, newName)
	time.Sleep(2 * time.Second)
}
