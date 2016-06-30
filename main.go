package main

import (
	"fmt"
	"math/rand"
	"time"
)

func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func main() {
	var hiddenNum int = random(1, 10)
	var guessNum int
	var amt int = 1
	fmt.Println("Please enter a number 1-10... ")
	fmt.Scan(&guessNum)
	for ; guessNum != hiddenNum; fmt.Scan(&guessNum) {
		amt++
		if guessNum < hiddenNum {
			fmt.Printf("Too Low! Guess Higher... \n")
		} else {
			fmt.Printf("Too High! Lower... \n")
		}
	}
	if guessNum == hiddenNum {
		fmt.Printf("That's right! %d was the number! \n", hiddenNum)
		fmt.Printf("It took you %d guesses to get it correctly! \n", amt)
	}
	var play string
	fmt.Println("Would you like to play again? Y/N")
	fmt.Scan(&play)
	if play == "Y" || play == "y" {
		main()
	} else {
		if play == "n" || play == "N" {
			fmt.Println("Goodbye!")
		} else {
			fmt.Println("Not valid response. You don't deserve to play again...")
		}
	}
}
