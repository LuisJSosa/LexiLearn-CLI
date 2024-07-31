# LexiLearn Configuration Guide

This guide provides detailed information on how to configure the LexiLearn CLI application. While LexiLearn is designed to work out-of-the-box, you may want to customize certain aspects of the application to better suit your needs.

## Table of Contents

1. [Overview](#overview)
2. [Configuration File](#configuration-file)
3. [Deck Directory](#deck-directory)
4. [Environment Variables](#environment-variables)
5. [Customizing the Application](#customizing-the-application)
    - [Changing the Review Algorithm](#changing-the-review-algorithm)
    - [Customizing the CLI Prompts](#customizing-the-cli-prompts)
6. [Troubleshooting Configuration Issues](#troubleshooting-configuration-issues)

## Overview

LexiLearn uses a simple configuration setup to manage different aspects of the application, such as the directory where decks are stored and various runtime settings. This guide will help you understand and customize these configurations.

## Configuration File

Currently, LexiLearn does not use a dedicated configuration file. Instead, configuration is managed through the code and environment variables. Future versions of LexiLearn may include a configuration file for more advanced customization.

## Deck Directory

By default, LexiLearn stores all vocabulary decks in the `decks` directory. You can change the location of this directory by modifying the `decksDir` variable in the `main.go` file.

### Example

1. Open the `main.go` file.
2. Find the line that defines the `decksDir` variable:
   ```go
   var decksDir = "decks"
   ```
3. Change the value of `decksDir` to your desired directory path:
   ```go
   var decksDir = "/path/to/your/custom/decks/directory"
   ```

## Environment Variables

LexiLearn allows certain configurations to be set through environment variables. These can be used to override default settings without modifying the source code.

### Supported Environment Variables

- `LEXILEARN_DECK_DIR`: Sets the directory where decks are stored.

### Example

To set an environment variable in your terminal session:

#### On Linux/macOS:
```sh
export LEXILEARN_DECK_DIR="/path/to/your/custom/decks/directory"
```

#### On Windows:
```cmd
set LEXILEARN_DECK_DIR=C:pathtoyourcustomdecksdirectory
```

In the `main.go` file, modify the initialization of the `decksDir` variable to use the environment variable if set:

```go
var decksDir = getEnv("LEXILEARN_DECK_DIR", "decks")

func getEnv(key, fallback string) string {
if value, exists := os.LookupEnv(key); exists {
return value
}
return fallback
}
```

## Customizing the Application

You can customize LexiLearn by modifying the source code. Here are a few common customizations:

### Changing the Review Algorithm

The review algorithm determines when words should be reviewed based on user input. You can modify this algorithm in the `calculateNextReview` function in the `main.go` file.

#### Example

1. Open the `main.go` file.
2. Find the `calculateNextReview` function:
   ```go
   func calculateNextReview(confidence int) time.Time {
   switch confidence {
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
   ```
3. Modify the time intervals to suit your needs.

### Customizing the CLI Prompts

You can customize the CLI prompts to change how the application interacts with users.

#### Example

1. Open the `main.go` file.
2. Find the sections of the code where prompts are displayed. For example, the main menu prompt:
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
3. Modify the prompts as needed. For example, you can change the text or add new options.

## Troubleshooting Configuration Issues

If you encounter issues with the configuration, consider the following steps:

- **Environment Variable Not Recognized**: Ensure the environment variable is set correctly and that your terminal session recognizes it. You can check the variable by running `echo $VARIABLE_NAME` on Linux/macOS or `echo %VARIABLE_NAME%` on Windows.
- **Custom Directory Not Working**: Ensure the custom directory exists and that LexiLearn has the necessary permissions to read and write to it. You can create the directory manually if it does not exist.
- **Unexpected Behavior After Customization**: If you experience issues after modifying the code, revert to the original code and make changes incrementally to identify the source of the problem.

---

By following this configuration guide, you should be able to customize LexiLearn to better suit your needs. For further assistance, refer to the [USAGE.md](./USAGE.md) and [TROUBLESHOOTING.md](./TROUBLESHOOTING.md) files.
