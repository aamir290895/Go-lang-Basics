package main

import (
	"fmt"
	"math/rand"
	"time"
)

func rollDice() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(6) + 1
}

func playGame() string {
	firstRoll := rollDice()
	secondRoll := rollDice()
	sum := firstRoll + secondRoll

	fmt.Printf("You rolled %d and %d. Total: %d\n", firstRoll, secondRoll, sum)

	switch sum {
	case 7, 11:
		return "You win!"
	case 2, 3, 12:
		return "You lose!"
	default:
		return "Roll again."
	}
}

// func main() {
// 	fmt.Println("Welcome to the Dice Game!")

// 	var playAgain string

// 	for {
// 		result := playGame()
// 		fmt.Println(result)

// 		fmt.Print("Do you want to play again? (yes/no): ")
// 		fmt.Scanln(&playAgain)

// 		if playAgain != "yes" {
// 			fmt.Println("Thanks for playing! Goodbye.")
// 			break
// 		}
// 	}
// }
