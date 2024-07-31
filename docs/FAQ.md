# LexiLearn FAQ

This document addresses some of the frequently asked questions about the LexiLearn CLI application.

## Table of Contents

1. [General Questions](#general-questions)
    - [What is LexiLearn?](#what-is-lexilearn)
    - [Who is LexiLearn for?](#who-is-lexilearn-for)
2. [Installation](#installation)
    - [How do I install LexiLearn?](#how-do-i-install-lexilearn)
3. [Usage](#usage)
    - [How do I start LexiLearn?](#how-do-i-start-lexilearn)
    - [How do I add a new vocabulary word?](#how-do-i-add-a-new-vocabulary-word)
    - [How do I list all vocabulary words?](#how-do-i-list-all-vocabulary-words)
    - [How do I review vocabulary words?](#how-do-i-review-vocabulary-words)
4. [Deck Management](#deck-management)
    - [How do I create a new deck?](#how-do-i-create-a-new-deck)
    - [How do I select a deck?](#how-do-i-select-a-deck)
    - [How do I rename a deck?](#how-do-i-rename-a-deck)
    - [How do I delete a deck?](#how-do-i-delete-a-deck)
5. [Troubleshooting](#troubleshooting)
    - [What should I do if I encounter an error during installation?](#what-should-i-do-if-i-encounter-an-error-during-installation)
    - [What should I do if LexiLearn is not working as expected?](#what-should-i-do-if-lexilearn-is-not-working-as-expected)

## General Questions

### What is LexiLearn?

LexiLearn is a command-line interface (CLI) application designed to help users manage and learn vocabulary words. It allows users to add, list, and review vocabulary words interactively.

### Who is LexiLearn for?

LexiLearn is for anyone looking to enhance their vocabulary, including students, language learners, and anyone interested in expanding their word knowledge.

## Installation

### How do I install LexiLearn?

To install LexiLearn, follow the instructions provided in the [INSTALL.md](INSTALL.md) file. It includes steps for cloning the repository, installing Go, and building the application.

## Usage

### How do I start LexiLearn?

To start LexiLearn, navigate to the project directory in your terminal and run the following command:

```sh
./lexilearn
```

### How do I add a new vocabulary word?

To add a new vocabulary word, choose option 1 from the main menu. You will be prompted to enter the term and its definition.

### How do I list all vocabulary words?

To list all vocabulary words stored in the application, choose option 2 from the main menu. The application will display all the terms and their definitions.

### How do I review vocabulary words?

To review vocabulary words, choose option 3 from the main menu. The application will prompt you to recall the definition of each word and then display the correct definition.

## Deck Management

### How do I create a new deck?

To create a new deck, choose option 4 from the main menu to access the Manage Decks menu. Then, select option 2 to create a new deck and enter the name of the new deck.

### How do I select a deck?

To select an existing deck, choose option 4 from the main menu to access the Manage Decks menu. Then, select option 3 and enter the number of the deck you want to select.

### How do I rename a deck?

To rename an existing deck, choose option 4 from the main menu to access the Manage Decks menu. Then, select option 4, enter the number of the deck you want to rename, and provide the new name.

### How do I delete a deck?

To delete an existing deck, choose option 4 from the main menu to access the Manage Decks menu. Then, select option 5 and enter the number of the deck you want to delete.

## Troubleshooting

### What should I do if I encounter an error during installation?

If you encounter an error during installation, ensure that Go is installed correctly and that your system's PATH environment variable includes the Go binary path. Refer to the [INSTALL.md](INSTALL.md) file for detailed installation steps.

### What should I do if LexiLearn is not working as expected?

If LexiLearn is not working as expected, try the following steps:
- Ensure you are using the latest version of the application by pulling the latest changes from the repository.
- Run `go mod tidy` to ensure all dependencies are correctly installed.
- Check for any error messages in the terminal and refer to the [Go documentation](https://golang.org/doc/) for troubleshooting tips.

If you continue to face issues, you can seek help from the Go community or refer to the project's documentation.

---

By referring to this FAQ, you should be able to resolve common issues and understand how to use LexiLearn effectively. For further assistance, consult the [USAGE.md](./USAGE.md) and [INSTALL.md](INSTALL.md) files.
