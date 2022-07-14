package main

import (
	"errors"
	"fmt"
	"log"
)

type Rectangle struct {
	width  int
	length int
}

func (rect Rectangle) area() int {
	len := rect.length
	wid := rect.width
	r := Rectangle{wid, len}
	if rect.width < 0 || rect.length < 0 {
		length, width, err := r.error()
		if err != nil {
			fmt.Println(length, width)
			log.Fatal("You entered : width = ", width, " length = ", length, " Please enter positive numbers --->>>> ", err)
		}
	}
	area := rect.width * rect.length
	return area
}

func (rect Rectangle) perimeter() int {
	len := rect.length
	wid := rect.width
	r := Rectangle{wid, len}
	if rect.width < 0 || rect.length < 0 {
		length, width, err := r.error()
		if err != nil {
			fmt.Println(length, width)
			log.Fatal("You entered : width =", width, " length = ", length, " Please enter positive numbers --->>>> ", err)
		}
	}
	perimeter := (rect.width * 2) + (rect.length * 2)

	return perimeter

}

func (rect *Rectangle) error() (int, int, error) {
	var errInvalid = errors.New("passing invalid arguments")
	return rect.length, rect.width, errInvalid
}

func main() {
	r1 := Rectangle{3, -3}
	// fmt.Println(r1.perimeter())
	fmt.Println(r1.area())
}
