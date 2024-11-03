package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
)

func getRandomWord() string {
	words := []string{"python", "hangman", "challenge", "programming", "developer"}
	return words[rand.Intn(len(words))]
}

func displayWord(word string, guessedLetters map[rune]bool) string {
	display := ""
	for _, letter := range word {
		if guessedLetters[letter] {
			display += string(letter) + " "
		} else {
			display += "_ "
		}
	}
	return display
}

func main() {
	word := getRandomWord()
	guessedLetters := make(map[rune]bool)
	attempts := 6

	fmt.Println("Welcome to Hangman!")
	fmt.Println(displayWord(word, guessedLetters))

	reader := bufio.NewReader(os.Stdin)

	for attempts > 0 {
		fmt.Print("Guess a letter: ")
		input, _ := reader.ReadString('\n')
		guess := strings.TrimSpace(strings.ToLower(input))

		if len(guess) != 1 {
			fmt.Println("Please enter a single letter.")
			continue
		}

		letter := rune(guess[0])

		if guessedLetters[letter] {
			fmt.Println("You already guessed that letter.")
		} else if strings.ContainsRune(word, letter) {
			guessedLetters[letter] = true
			fmt.Println("Good guess!")
		} else {
			guessedLetters[letter] = true
			attempts--
			fmt.Printf("Wrong guess. You have %d attempts left.\n", attempts)
		}

		currentDisplay := displayWord(word, guessedLetters)
		fmt.Println(currentDisplay)

		if !strings.Contains(currentDisplay, "_") {
			fmt.Println("Congratulations! You've guessed the word!")
			break
		}
	}

	if attempts == 0 {
		fmt.Printf("Sorry, you've run out of attempts. The word was '%s'.\n", word)
	}
}
