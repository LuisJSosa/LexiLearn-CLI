# Go Vocabulary CLI

This is a CLI application for managing vocabulary words. The application allows you to add and list vocabulary words.

## Commands

- `add <term> <definition>`: Adds a new vocabulary word with the given term and definition.
- `list`: Lists all vocabulary words.

## Setup

1. Install [Go](https://golang.org/doc/install).
2. Clone the repository.
3. Run `go mod tidy` to install dependencies.

## Usage

To run the application:
```sh
go run main.go 
```
To add a word:
```sh
go run cmd/main.go add "term" "definition"
```
#### go.mod:
```go
module GoVocab-CLI

go 1.20