# LexiLearn CLI

LexiLearn is a command-line interface (CLI) application designed to help users manage and learn vocabulary words. The application allows users to add, list, and review vocabulary words, and it is built using the Go programming language.

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Setup](#setup)
4. [Usage](#usage)
    - [Main Menu](#main-menu)
    - [Adding Vocabulary Words](#adding-vocabulary-words)
    - [Listing Vocabulary Words](#listing-vocabulary-words)
    - [Reviewing Vocabulary Words](#reviewing-vocabulary-words)
    - [Managing Decks](#managing-decks)
        - [Listing Decks](#listing-decks)
        - [Creating a New Deck](#creating-a-new-deck)
        - [Selecting a Deck](#selecting-a-deck)
        - [Renaming a Deck](#renaming-a-deck)
        - [Deleting a Deck](#deleting-a-deck)
    - [Exiting the Application](#exiting-the-application)
5. [Development](#development)
6. [Testing](#testing)
7. [Contributing](#contributing)
8. [License](#license)
9. [Installation Instructions](#installation-instructions)


## Introduction

LexiLearn is a simple and effective tool for anyone looking to enhance their vocabulary. By using this CLI application, users can add new vocabulary words, list all stored words, and review them in an interactive manner.

## Features

- **Add Vocabulary Words**: Users can add new vocabulary words along with their definitions.
- **List Vocabulary Words**: Users can list all the vocabulary words stored in the application.
- **Review Vocabulary Words**: Users can review the vocabulary words interactively to reinforce their learning.
- **Manage Decks**: Users can manage multiple vocabulary decks, including listing, creating, selecting, renaming, and deleting decks.

## Setup

To set up the LexiLearn CLI application, follow these steps:

### Step 1: Clone the Repository

First, you need to clone the repository and navigate to the LexiLearn directory:

```sh
git clone https://github.com/LuisJSosa/LexiLearn-CLI.git
cd LexiLearn-CLI
```

### Step 2: Install Go and Necessary Tools

Download and install Go from [the official website](https://golang.org/dl/).

### Step 3: Install Dependencies

Run the following command to install the necessary dependencies:

```sh
go mod tidy
```

## Usage

To run the LexiLearn CLI application, use the following command:

```sh
go run menu.go
```

### Main Menu

When you run LexiLearn, you will see the main menu:

```
======================================
         Welcome to LexiLearn         
======================================
Current Deck: Default
Please choose an option:
1. Add a new vocabulary word
2. List all vocabulary words
3. Review vocabulary words
4. Manage decks
5. Exit
   ======================================
   Enter your choice:
   ```

### Adding Vocabulary Words

To add a new vocabulary word, select option 1 from the main menu. You will be prompted to enter the term and its definition.

Example:

```
You chose to add a new vocabulary word, in Deck: Default.
No deck selected, progress saved in Default deck.
Enter the term: apple
Enter the definition: a fruit

Added: apple - a fruit
```

### Listing Vocabulary Words

To list all vocabulary words, select option 2 from the main menu. The vocabulary list will be displayed along with the next review time.

Example:

```
======================================
          Vocabulary List             
======================================
1. apple: a fruit (Next review: 01-02-2022 03:04 PM)
2. book: a collection of pages (Next review: 01-02-2022 03:04 PM)

Enter the number of the word to delete or press Enter to return to the main menu:
```

### Reviewing Vocabulary Words

To review vocabulary words, select option 3 from the main menu. You will be prompted to think of the definition of each word and then rate your confidence.

Example:

```
Think of the definition of 'apple'. Press Enter to see the correct definition.

The correct definition is: a fruit
How confident do you feel about this word? (1-5): 4

Next review for 'apple' scheduled.
======================================
```

### Managing Decks

To manage your decks, select option 4 from the main menu. You can list, create, select, rename, and delete decks.

#### Listing Decks

Example:

```
======================================
            Manage Decks              
======================================
1. List decks
2. Create a new deck
3. Select a deck
4. Rename a deck
5. Delete a deck
6. Back to main menu
   ======================================
   Enter your choice: 1

Available Decks:
1. Default
2. Spanish
   ```

#### Creating a New Deck

Example:

```
======================================
            Manage Decks              
======================================
1. List decks
2. Create a new deck
3. Select a deck
4. Rename a deck
5. Delete a deck
6. Back to main menu
   ======================================
   Enter your choice: 2

You chose to create a new deck.
Enter the name of the new deck: French

Deck 'French' created successfully!
```

#### Selecting a Deck

Example:

```
======================================
            Manage Decks              
======================================
1. List decks
2. Create a new deck
3. Select a deck
4. Rename a deck
5. Delete a deck
6. Back to main menu
   ======================================
   Enter your choice: 3

Available Decks:
1. Default
2. French
3. Spanish

Enter the number of the deck to select: 2

Deck 'French' selected successfully!
```

#### Renaming a Deck

Example:

```
======================================
            Manage Decks              
======================================
1. List decks
2. Create a new deck
3. Select a deck
4. Rename a deck
5. Delete a deck
6. Back to main menu
   ======================================
   Enter your choice: 4

Available Decks:
1. Default
2. French
3. Spanish

Enter the number of the deck to rename: 2
Enter the new name for the deck: German

Deck 'French' renamed to 'German' successfully!
```

#### Deleting a Deck

Example:

```
======================================
            Manage Decks              
======================================
1. List decks
2. Create a new deck
3. Select a deck
4. Rename a deck
5. Delete a deck
6. Back to main menu
   ======================================
   Enter your choice: 5

Available Decks:
1. Default
2. German
3. Spanish

Enter the number of the deck to delete: 2

Deck 'German' deleted successfully!
```

### Exiting the Application

To exit LexiLearn, select option 5 from the main menu:

```
Exiting LexiLearn... Goodbye!
```

## Development

The project structure is as follows:

```
lexilearn
├── decks
│   └── deck1.json
│   
├── docs
│   └── ARCHITECTURE.md
│   └── CONFIG.md
│   └── FAQ.md
│   └── INSTALL.md
│   └── TROUBLESHOOTING.md
│   └── USAGE.md
│  
├── CODE_OF_CONDUCT.md
├── CONTRIBUTING.md
│  
├── go.mod
│   └── go.sum
│   
├── main.go
├── main_test.go
├── renovate.json
├── LICENSE
│  
├── External Libraries
│   └── Go Modules
│      └── etc..
│   └── Go SDK 1.23rc1
│      └── etc..
│   
└── README.md
```

- `cmd/menu.go`: The entry point for the application containing the main menu logic.
- `internal/vocabulary.go`: Contains the data structures and functions for managing vocabulary words.
- `go.mod`: The Go module file.
- `go.sum`: The Go dependencies file.
- `README.md`: The project documentation.

## Testing

To run the tests for the LexiLearn CLI application, use the following command:

```sh
go test
```

This will execute the unit tests and ensure that all functionalities work as expected.

## Contributing

Contributions are welcome! If you would like to contribute to this project, please follow these steps:

1. Fork the repository.
2. Create a new branch for your feature or bug fix.
3. Commit your changes with a descriptive commit message.
4. Push your changes to your fork.
5. Create a pull request to the main repository.

Please make sure your code adheres to the project's coding standards and passes all tests.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Installation Instructions

For detailed installation instructions on running LexiLearn in a Kubernetes environment, please refer to the [Running LexiLearn in Kubernetes](docs/running_app_in_kubernetes.md) guide.




