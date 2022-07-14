package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func generateNum() [4]int { // generating the random number array with size of 4.
	var numArr [4]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		numArr[i] = rand.Intn(8) + 1
	}
	return numArr
}

func checkGuess(computer []int, user []int) (int, int) { // this is a func of handles with counters.
	greenCounter := 0  // green counter means, one of the number of user entered is on correct place.
	yellowCounter := 0 // yellow counter means, one of the number of user entered is in the array but on wrong place.
	for i := 0; i < len(computer); i++ {
		if computer[i] == user[i] {
			greenCounter++ // controling for green counter.
		}
		a := make([]int, 0)
		a = append(a, computer[:i]...)
		c := append(a, computer[i+1:]...)
		for j := 0; j < len(c); j++ { // controling for other indexes.
			if user[i] == c[j] {
				yellowCounter++
			}
		}
	}
	return greenCounter, yellowCounter
}

func main() {
	compArr := generateNum() // calling generateNum() func.
	fmt.Println(compArr)
	var userArr [4]int
	prevGuesses := make([]int, 0)              // this is an array for handling previous guesses.
	prevGreens := make([]int, 0)               // this is an array for handling previous greenCounters.
	prevYellows := make([]int, 0)              // this is an array for handling previous yellowCounters.
	for !reflect.DeepEqual(compArr, userArr) { // equality control for slices.
		for i := 0; i < 4; i++ {
			fmt.Println("ENTER A NUMBER")
			fmt.Scanln(&userArr[i]) // taking numbers from user.
		}
		prevGuesses = append(prevGuesses, userArr[:]...) // appending the guess that user made.

		greenCounter, yellowCounter := checkGuess(compArr[:], userArr[:])           //handling the return values from checkGuess() func.
		prevGreens = append(prevGreens, greenCounter)                               // appending the greenCounter.
		prevYellows = append(prevYellows, yellowCounter)                            // appending the yellowCounter.
		for i, counter := 0, 0; i < len(prevGuesses); i, counter = i+4, counter+1 { //printing the previous guesses and counters.
			fmt.Printf("Your %d. guess is : %v ---->>> Green Counter is = %d  Yellow Counter is = %d\n", counter+1, prevGuesses[i:i+4], prevGreens[counter], prevYellows[counter])
		}

		if greenCounter == 4 {
			fmt.Printf("You guessed correctly in %d tries! Well Done!\n", len(prevGuesses)/4) // finish print.
			break
		}
	}
}
