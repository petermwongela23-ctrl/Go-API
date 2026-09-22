package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

// playRound runs one game and returns (won, attempts).
func playRound(reader *bufio.Reader) (bool, int) {
	const maxNumber = 100
	const maxAttempts = 7
	secret := rand.Intn(maxNumber) + 1
	attempts := 0

	fmt.Printf("\nI'm thinking of a number between 1 and %d.\n", maxNumber)
	fmt.Printf("You have %d guesses. Good luck!\n", maxAttempts)

	for attempts < maxAttempts {
		fmt.Print("\nEnter your guess (or 'q' to quit): ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		if input == "q" || input == "Q" {
			fmt.Printf("No worries! The number was %d.\n", secret)
			return false, attempts
		}

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("That's not a valid number. Try again.")
			continue
		}

		if guess < 1 || guess > maxNumber {
			fmt.Printf("Please guess between 1 and %d.\n", maxNumber)
			continue
		}

		attempts++

		switch {
		case guess < secret:
			fmt.Println("Too low! Aim higher.")
		case guess > secret:
			fmt.Println("Too high! Aim lower.")
		default:
			fmt.Printf("\nYou got it! It took %d attempt(s).\n", attempts)
			return true, attempts
		}

		remaining := maxAttempts - attempts
		if remaining > 0 {
			fmt.Printf("Guesses left: %d\n", remaining)
		}
	}

	fmt.Printf("\nOut of guesses! The number was %d. Better luck next time!\n", secret)
	return false, attempts
}

// askPlayAgain returns true if the player wants another round.
func askPlayAgain(reader *bufio.Reader) bool {
	fmt.Print("\nPlay again? (y/n): ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(strings.ToLower(input))
	return input == "y" || input == "yes"
}

func main() {
	rand.Seed(time.Now().UnixNano())
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Guess the Number!")

	bestScore := 0 // 0 means "no wins yet"

	for {
		won, attempts := playRound(reader)

		if won {
			if bestScore == 0 || attempts < bestScore {
				bestScore = attempts
				fmt.Println("That's a new best score!")
			}
			fmt.Printf("Best score so far: %d guess(es).\n", bestScore)
		}

		if !askPlayAgain(reader) {
			fmt.Println("\nThanks for playing. Goodbye!")
			break
		}
	}
}
