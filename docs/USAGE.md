# LexiLearn Usage Guide

This guide will help you understand how to use the LexiLearn CLI application effectively. It includes detailed instructions and examples for each feature of the application.

## Table of Contents

1. [Introduction](#introduction)
2. [Running the Application](#running-the-application)
3. [Main Menu](#main-menu)
4. [Adding Vocabulary Words](#adding-vocabulary-words)
5. [Listing Vocabulary Words](#listing-vocabulary-words)
6. [Reviewing Vocabulary Words](#reviewing-vocabulary-words)
7. [Managing Decks](#managing-decks)
    - [Listing Decks](#listing-decks)
    - [Creating a New Deck](#creating-a-new-deck)
    - [Selecting a Deck](#selecting-a-deck)
    - [Renaming a Deck](#renaming-a-deck)
    - [Deleting a Deck](#deleting-a-deck)
8. [Exiting the Application](#exiting-the-application)

## Introduction

LexiLearn is a command-line application designed to help you manage and learn vocabulary words through an interactive interface. You can add new words, review them using spaced repetition, and manage multiple vocabulary decks.

## Running the Application

To start the LexiLearn CLI application, navigate to the project directory and run the following command:

```sh
./lexilearn
```

You will see the main menu of the application.

## Main Menu

The main menu provides several options to interact with the LexiLearn application:

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

## Adding Vocabulary Words

To add a new vocabulary word, select option 1 from the main menu. You will be prompted to enter the term and its definition.

### Example

```
You chose to add a new vocabulary word, in Deck: Default.
No deck selected, progress saved in Default deck.
Enter the term: apple
Enter the definition: a fruit

Added: apple - a fruit
```

## Listing Vocabulary Words

To list all vocabulary words, select option 2 from the main menu. The vocabulary list will be displayed along with the next review time.

### Example

```
======================================
          Vocabulary List             
======================================
1. apple: a fruit (Next review: 01-02-2022 03:04 PM)
2. book: a collection of pages (Next review: 01-02-2022 03:04 PM)

Enter the number of the word to delete or press Enter to return to the main menu:
```

## Reviewing Vocabulary Words

To review vocabulary words, select option 3 from the main menu. You will be prompted to think of the definition of each word and then rate your confidence.

### Example

```
Think of the definition of 'apple'. Press Enter to see the correct definition.

The correct definition is: a fruit
How confident do you feel about this word? (1-5): 4

Next review for 'apple' scheduled.
======================================
```

## Managing Decks

To manage your decks, select option 4 from the main menu. You can list, create, select, rename, and delete decks.

### Listing Decks

To list all available decks, choose option 1 under the Manage Decks menu.

#### Example

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

### Creating a New Deck

To create a new deck, choose option 2 under the Manage Decks menu. You will be prompted to enter the name of the new deck.

#### Example

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

### Selecting a Deck

To select an existing deck, choose option 3 under the Manage Decks menu. You will be prompted to enter the number of the deck to select.

#### Example

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

### Renaming a Deck

To rename an existing deck, choose option 4 under the Manage Decks menu. You will be prompted to enter the number of the deck to rename and the new name.

#### Example

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

### Deleting a Deck

To delete an existing deck, choose option 5 under the Manage Decks menu. You will be prompted to enter the number of the deck to delete.

#### Example

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

## Exiting the Application

To exit LexiLearn, select option 5 from the main menu:

```
Exiting LexiLearn... Goodbye!
```

## Summary

This usage guide has provided you with detailed instructions and examples to use LexiLearn effectively. You can add, review, and manage vocabulary words and decks easily. Enjoy learning your vocabulary with LexiLearn!
