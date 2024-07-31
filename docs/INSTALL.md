# Installation Guide for LexiLearn

This guide will help you set up and install the LexiLearn CLI application on your system.

## Table of Contents

1. [Prerequisites](#prerequisites)
2. [Clone the Repository](#clone-the-repository)
3. [Install Go](#install-go)
4. [Build the Application](#build-the-application)
5. [Run LexiLearn](#run-lexilearn)
6. [Troubleshooting](#troubleshooting)

## Prerequisites

Before you begin, ensure you have the following installed on your system:

- **Git**: A version control system to clone the repository.
- **Go**: The Go programming language. You can download it from [Go's official website](https://golang.org/dl/).

## Clone the Repository

1. Open your terminal or command prompt.
2. Clone the LexiLearn repository using the following command:

   ```sh
   git clone https://github.com/LuisJSosa/GoVocab-CLI.git
   ```

3. Navigate to the LexiLearn directory:

   ```sh
   cd GoVocab-CLI
   ```

## Install Go

1. Download and install Go from [the official website](https://golang.org/dl/).

2. Follow the installation instructions provided on the website for your specific operating system.

3. Verify the installation by running the following command in your terminal or command prompt:

   ```sh
   go version
   ```

   You should see output indicating the installed Go version.

## Build the Application

1. Ensure you are in the LexiLearn project directory:

   ```sh
   cd GoVocab-CLI
   ```

2. Install the necessary dependencies:

   ```sh
   go mod tidy
   ```

3. Build the application:

   ```sh
   go build -o lexilearn main.go
   ```

   This command will generate an executable file named `lexilearn` in the project directory.

## Run LexiLearn

To start the LexiLearn CLI application, run the following command:

```sh
./lexilearn
```

You should see the main menu of the LexiLearn application:

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

## Troubleshooting

If you encounter any issues during installation or setup, consider the following steps:

- **Go installation issues**: Ensure that Go is installed correctly and that your system's PATH environment variable includes the Go binary path.
- **Dependency issues**: Run `go mod tidy` to ensure all dependencies are correctly installed.
- **Build issues**: Verify that you are in the correct directory and that all files are present.

If you continue to face issues, please refer to the [Go documentation](https://golang.org/doc/) or seek help from the Go community.

---

By following these steps, you should have the LexiLearn CLI application installed and running on your system. Enjoy learning your vocabulary with LexiLearn!
