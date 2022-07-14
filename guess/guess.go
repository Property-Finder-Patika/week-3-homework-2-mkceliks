package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func generateNum() [4]int {
	var numArr [4]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		numArr[i] = rand.Intn(8) + 1
	}
	return numArr
}

func checkGuess(computer []int, user []int) (int, int) {
	greenCounter := 0
	yellowCounter := 0
	for i := 0; i < len(computer); i++ {
		if computer[i] == user[i] {
			greenCounter++
		}
		a := make([]int, 0)
		a = append(a, computer[:i]...)
		c := append(a, computer[i+1:]...)
		for j := 0; j < len(c); j++ {
			if user[i] == c[j] {
				yellowCounter++
			}
		}
	}
	return greenCounter, yellowCounter
}

func main() {
	compArr := generateNum()
	fmt.Println(compArr)
	var userArr [4]int
	prevGuesses := make([]int, 0)
	prevGreens := make([]int, 0)
	prevYellows := make([]int, 0)
	for !reflect.DeepEqual(compArr, userArr) {

		for i := 0; i < 4; i++ {
			fmt.Println("ENTER A NUMBER")
			fmt.Scanln(&userArr[i])
		}
		prevGuesses = append(prevGuesses, userArr[:]...)

		greenCounter, yellowCounter := checkGuess(compArr[:], userArr[:])
		prevGreens = append(prevGreens, greenCounter)
		prevYellows = append(prevYellows, yellowCounter)
		for i, counter := 0, 0; i < len(prevGuesses); i, counter = i+4, counter+1 {
			fmt.Printf("Your %d. guess is : %v ---->>> Green Counter is = %d  Yellow Counter is = %d\n", counter+1, prevGuesses[i:i+4], prevGreens[counter], prevYellows[counter])
		}

		if greenCounter == 4 {
			fmt.Printf("You guessed correctly in %d tries! Well Done!\n", len(prevGuesses)/4)
			break
		}
	}
}
