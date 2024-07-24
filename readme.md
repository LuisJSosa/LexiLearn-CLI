# LexiLearn CLI

LexiLearn is a command-line interface (CLI) application designed to help users manage and learn vocabulary words. The application allows users to add, list, and review vocabulary words, and it is built using the Go programming language.

## Table of Contents

1. [Introduction](#introduction)
2. [Features](#features)
3. [Setup](#setup)
4. [Usage](#usage)
5. [Commands](#commands)
6. [Development](#development)
7. [Testing](#testing)
8. [Contributing](#contributing)
9. [License](#license)

## Introduction

LexiLearn is a simple and effective tool for anyone looking to enhance their vocabulary. By using this CLI application, users can add new vocabulary words, list all stored words, and review them in an interactive manner.

## Features

- **Add Vocabulary Words**: Users can add new vocabulary words along with their definitions.
- **List Vocabulary Words**: Users can list all the vocabulary words stored in the application.
- **Review Vocabulary Words**: Users can review the vocabulary words interactively to reinforce their learning.

## Setup

To set up the LexiLearn CLI application, follow these steps:

1. **Install Go**: Ensure you have Go installed on your system. You can download it from [Go's official website](https://golang.org/dl/).

2. **Clone the Repository**:
    ```sh
    git clone https://github.com/LuisJSosa/GoVocab-CLI.git
    cd GoVocab-CLI
    ```

3. **Install Dependencies**:
    ```sh
    go mod tidy
    ```

## Usage

To run the LexiLearn CLI application, use the following command:

```sh
go run menu.go
```

Upon running the application, you will see the main menu with the following options:

```
======================================
         Welcome to LexiLearn         
======================================
Please choose an option:
1. Add a new vocabulary word
2. List all vocabulary words
3. Review vocabulary words
4. Exit
======================================
Enter your choice: 
```

## Commands

### Add a Vocabulary Word

To add a new vocabulary word, choose option 1. You will be prompted to enter the term and its definition.

### List Vocabulary Words

To list all the vocabulary words stored in the application, choose option 2. The application will display all the terms and their definitions.

### Review Vocabulary Words

To review the vocabulary words, choose option 3. The application will prompt you to recall the definition of each word and then display the correct definition.

### Exit

To exit the application, choose option 4.

## Development

The project structure is as follows:

```
lexilearn
├── cmd
│   └── menu.go
├── internal
│   └── vocabulary.go
├── go.mod
├── go.sum
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
