# LexiLearn Architecture Guide

This guide provides a detailed overview of the architecture of the LexiLearn CLI application. Understanding the architecture will help you navigate the codebase, make modifications, and extend the functionality of the application.

## Table of Contents

1. [Overview](#overview)
2. [Project Structure](#project-structure)
3. [Main Components](#main-components)
    - [main.go](#maingo)
    - [main_test.go](#maintestgo)
4. [Data Structures](#data-structures)
    - [Word](#word)
    - [Deck](#deck)
5. [Application Flow](#application-flow)
6. [Function Breakdown](#function-breakdown)
7. [Dependencies](#dependencies)
8. [Future Improvements](#future-improvements)

## Overview

LexiLearn is a command-line interface (CLI) application designed to help users manage and learn vocabulary words. It is built using the Go programming language and follows a simple yet effective architecture to handle vocabulary management and review.

## Project Structure

The LexiLearn project is organized as follows:

```
lexilearn
├── decks
│   └── deck1.json
│   
├── go.mod
│   └── go.sum
│   
├── main.go
├── main_test.go
├── renovate.json
├── LICENSE
├── External Libraries
│   └── Go Modules
│      └── etc..
│   └── Go SDK 1.23rc1
│      └── etc..
│   
└── README.md
```

### Directories and Files

- **decks/**: Directory where vocabulary decks are stored as JSON files.
- **go.mod** and **go.sum**: Go module files for dependency management.
- **main.go**: The main entry point of the application containing the core logic.
- **main_test.go**: Contains unit tests for the application.
- **renovate.json**: Configuration file for dependency management using Renovate.
- **LICENSE**: License file for the project.
- **README.md**: Main documentation file providing an overview and usage instructions.

## Main Components

### main.go

The `main.go` file contains the main logic of the LexiLearn application. It includes functions for displaying menus, managing vocabulary words, reviewing words, and handling deck management.

### main_test.go

The `main_test.go` file contains unit tests to ensure the functionality of the LexiLearn application. It tests various functions and ensures that the application behaves as expected.

## Data Structures

### Word

The `Word` struct represents a vocabulary word and its associated data.

```go
type Word struct {
Term         string
Definition   string
CorrectCount int
LastReview   time.Time
NextReview   time.Time
}
```

### Deck

The `Deck` struct is not explicitly defined in the current version of LexiLearn. Instead, decks are managed through JSON files in the `decks/` directory, with each file representing a deck.

## Application Flow

1. **Initialization**: The application starts by setting up the necessary directories and loading the vocabulary deck from the JSON file.
2. **Main Menu**: The user is presented with a main menu to choose actions such as adding words, listing words, reviewing words, or managing decks.
3. **Adding Words**: Users can add new vocabulary words by entering the term and definition.
4. **Listing Words**: The application lists all vocabulary words stored in the current deck.
5. **Reviewing Words**: The application prompts users to review words based on a spaced repetition algorithm.
6. **Managing Decks**: Users can create, select, rename, and delete decks.

## Function Breakdown

### displayMenu

Displays the main menu options to the user.

```go
func displayMenu() {
fmt.Println("======================================")
fmt.Println("         Welcome to LexiLearn         ")
fmt.Println("======================================")
fmt.Println(color.GreenString("Current Deck: ") + color.MagentaString(deckName))
fmt.Println("Please choose an option:")
fmt.Println("1. Add a new vocabulary word")
fmt.Println("2. List all vocabulary words")
fmt.Println("3. Review vocabulary words")
fmt.Println("4. Manage decks")
fmt.Println("5. Exit")
fmt.Println("======================================")
fmt.Print("Enter your choice: ")
}
```

### addWord

Prompts the user to add a new vocabulary word and its definition.

```go
func addWord() {
reader := bufio.NewReader(os.Stdin)
clearScreen()
if deckName != "Default" {
color.Cyan("You chose to add a new vocabulary word, in Deck: %s.", deckName)
} else {
color.Cyan("You chose to add a new vocabulary word.")
color.Yellow("No deck selected, progress saved in Default deck.")
}

    fmt.Print("Enter the term: ")
    term, _ := reader.ReadString('n')
    term = strings.TrimSpace(term) // Remove trailing newline
    fmt.Print("Enter the definition: ")
    definition, _ := reader.ReadString('n')
    definition = strings.TrimSpace(definition) // Remove trailing newline

    // Validate inputs
    if term == "" || definition == "" {
        color.Red("Both term and definition are required.")
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

    color.Green("Added: %s - %sn", term, definition)
    saveVocabList() // Save after adding a word
    time.Sleep(2 * time.Second)
}
```

### listWords

Lists all vocabulary words in the current deck.

```go
func listWords() {
clearScreen()

    if len(vocabList) == 0 {
        color.Red("No vocabulary words found.")
        return
    }
    color.Blue("Vocabulary List:")
    for i, word := range vocabList {
        fmt.Printf("%d. %s: %s (Next review: %s)n", i+1, word.Term, word.Definition, word.NextReview.Format("01-02-2006 03:04 PM"))
    }
    fmt.Print("Enter the number of the word to delete or press Enter to return to the main menu: ")
    var choice int
    _, err := fmt.Scanf("%d", &choice)
    if err == nil && choice > 0 && choice <= len(vocabList) {
        deleteWord(choice - 1)
    }
    time.Sleep(5 * time.Second)
}
```

### reviewWords

Prompts the user to review vocabulary words based on a spaced repetition algorithm.

```go
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
            _, _ = reader.ReadString('n')

            color.Cyan("The correct definition is: %sn", word.Definition)

            var confidence int
            for {
                fmt.Print("How confident do you feel about this word? (1-5): ")
                input, err := reader.ReadString('n')
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
            color.Cyan("Next review for '%s' scheduled.n", word.Term)
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
            cont, _ := reader.ReadString('n')
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
            cont, _ := reader.ReadString('n')
            cont = strings.TrimSpace(cont)
            if strings.ToLower(cont) != "yes" {
                break
            }
        }
    }
}
```

## Dependencies

LexiLearn relies on a few external libraries to enhance its functionality:

- **fatih/color**: Used to add color to CLI output for better user experience.
- **encoding/json**: Used to encode and decode JSON data for storing and retrieving vocabulary decks.

Dependencies are managed using Go modules. Ensure you have the required dependencies by running:

```sh
go mod tidy
```

## Future Improvements

While LexiLearn is functional, there are several areas for potential improvement:

- **Configuration File**: Introduce a dedicated configuration file for easier customization.
- **User Interface**: Enhance the CLI user interface for a more interactive and user-friendly experience.
- **Advanced Review Algorithms**: Implement more sophisticated spaced repetition algorithms to optimize learning.
- **Web Interface**: Develop a web-based interface to complement the CLI application, providing more flexibility for users.

---

By understanding the architecture of LexiLearn, you can navigate the codebase more effectively and contribute to its development. For further assistance, refer to the [USAGE.md](./USAGE.md) and [TROUBLESHOOTING.md](./TROUBLESHOOTING.md) files.
