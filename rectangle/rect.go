package main

import (
	"errors"
	"fmt"
	"log"
)

type Rectangle struct { // struct of Rectangle
	width  int
	length int
}

func (rect Rectangle) area() int { // function of area calculator.
	len := rect.length
	wid := rect.width
	r := Rectangle{wid, len}
	if rect.width < 0 || rect.length < 0 { // this if statement is about Error-handling that handles the negative nums.
		length, width, err := r.error()
		if err != nil {
			fmt.Println(length, width)
			log.Fatal("You entered : width = ", width, " length = ", length, " Please enter positive numbers --->>>> ", err)
		}
	}
	area := rect.width * rect.length // calculating the area of rectangle and returning.
	return area
}

func (rect Rectangle) perimeter() int { // function of perimeter calculator.
	len := rect.length
	wid := rect.width
	r := Rectangle{wid, len}
	if rect.width < 0 || rect.length < 0 { // this if statement is about Error-handling that handles the negative nums.
		length, width, err := r.error()
		if err != nil {
			fmt.Println(length, width)
			log.Fatal("You entered : width =", width, " length = ", length, " Please enter positive numbers --->>>> ", err)
		}
	}
	perimeter := (rect.width * 2) + (rect.length * 2) // calculating the perimeter of rectangle and returning.

	return perimeter

}

func (rect *Rectangle) error() (int, int, error) { // error function.
	var errInvalid = errors.New("passing invalid arguments")
	return rect.length, rect.width, errInvalid
}

func main() {
	r1 := Rectangle{3, -3}
	fmt.Println(r1.perimeter()) // calling perimeter func with r1.
	fmt.Println(r1.area())      // calling area func with r1.
}
