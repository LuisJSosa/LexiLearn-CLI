# Troubleshooting Guide for LexiLearn

This guide provides detailed troubleshooting steps for common issues users might encounter while using the LexiLearn CLI application.

## Table of Contents

1. [General Issues](#general-issues)
    - [LexiLearn won't start](#lexilearn-wont-start)
    - [Commands not recognized](#commands-not-recognized)
2. [Installation Issues](#installation-issues)
    - [Go installation issues](#go-installation-issues)
    - [Dependency issues](#dependency-issues)
3. [Usage Issues](#usage-issues)
    - [Error adding a new vocabulary word](#error-adding-a-new-vocabulary-word)
    - [Error listing vocabulary words](#error-listing-vocabulary-words)
    - [Error reviewing vocabulary words](#error-reviewing-vocabulary-words)
4. [Deck Management Issues](#deck-management-issues)
    - [Error creating a new deck](#error-creating-a-new-deck)
    - [Error selecting a deck](#error-selecting-a-deck)
    - [Error renaming a deck](#error-renaming-a-deck)
    - [Error deleting a deck](#error-deleting-a-deck)

## General Issues

### LexiLearn won't start

**Solution**:
1. Ensure you have built the application by running `go build -o lexilearn main.go`.
2. Verify that you are in the correct directory by running `ls`. You should see the `lexilearn` executable listed.
3. Run the application using `./lexilearn`.

### Commands not recognized

**Solution**:
1. Ensure you are entering the correct commands as listed in the [Usage Guide](./USAGE.md).
2. Make sure you are in the correct directory where the `lexilearn` executable is located.
3. Verify that the executable has the correct permissions. You may need to run `chmod +x lexilearn`.

## Installation Issues

### Go installation issues

**Solution**:
1. Verify that Go is installed by running `go version`. You should see the installed version of Go.
2. Ensure your system's PATH environment variable includes the Go binary path. You can add it to your PATH by following the instructions in the [Go installation guide](https://golang.org/doc/install).

### Dependency issues

**Solution**:
1. Run `go mod tidy` to ensure all dependencies are correctly installed.
2. Check the `go.mod` and `go.sum` files for any missing or incorrect dependencies.

## Usage Issues

### Error adding a new vocabulary word

**Solution**:
1. Ensure you are following the correct format when adding a word. You should be prompted to enter the term and its definition.
2. Check for any error messages in the terminal that might indicate what went wrong.

### Error listing vocabulary words

**Solution**:
1. Ensure there are vocabulary words added to the deck. If the deck is empty, there will be no words to list.
2. Check for any error messages in the terminal that might indicate what went wrong.

### Error reviewing vocabulary words

**Solution**:
1. Ensure there are vocabulary words added to the deck. If the deck is empty, there will be no words to review.
2. Check for any error messages in the terminal that might indicate what went wrong.

## Deck Management Issues

### Error creating a new deck

**Solution**:
1. Ensure you are entering a valid name for the new deck. Deck names should not be empty and should not contain special characters.
2. Check for any error messages in the terminal that might indicate what went wrong.

### Error selecting a deck

**Solution**:
1. Ensure the deck you are trying to select exists. Use the "List decks" option to verify available decks.
2. Check for any error messages in the terminal that might indicate what went wrong.

### Error renaming a deck

**Solution**:
1. Ensure the deck you are trying to rename exists. Use the "List decks" option to verify available decks.
2. Ensure you are entering a valid new name for the deck. Deck names should not be empty and should not contain special characters.
3. Check for any error messages in the terminal that might indicate what went wrong.

### Error deleting a deck

**Solution**:
1. Ensure the deck you are trying to delete exists. Use the "List decks" option to verify available decks.
2. Be cautious when deleting decks, as this action cannot be undone.
3. Check for any error messages in the terminal that might indicate what went wrong.

---

By following these troubleshooting steps, you should be able to resolve common issues encountered while using LexiLearn. For further assistance, consult the [FAQ](./FAQ.md) and [USAGE.md](./USAGE.md) files.
