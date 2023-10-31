package main

import (
	"fmt"
	"math/rand"
)

func main() {
	target := rand.Intn(100) + 1
	fmt.Println("I've chosen a random number between 1 and 100.")
	fmt.Println("Can you guess it?")

	success := false
	for guesses := 0; guesses <= 10; guesses++ {
		fmt.Println("You have", 10-guesses, "guesses left.")

		fmt.Print("Make a guess: ")
		var input int
		fmt.Scan(&input)

		if input > target {
			fmt.Println("Your guess was high.")
		} else if input < target {
			fmt.Println("Your guess was low.")
		} else {
			success = true
			fmt.Println("Good job! You guessed it!")
			break
		}
	}

	if !success {
		fmt.Println("Sorry you didn't guess my number. It was:", target)
	}
}
