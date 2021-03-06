package main

import (
	"fmt"
)

type Plate struct {
	pasta     string
	vegetable string
}

type JustCook func(string)

func CookLunch(lunch Plate, do func(string)) {
	course := lunch.pasta + " with " + lunch.vegetable
	do(course)
}

func CookDinner(lunch Plate, do JustCook) {
	course := lunch.pasta + " with " + lunch.vegetable
	do(course)
}

func SimpleMenu(menu string) {
	fmt.Println("Main Course: ", menu)
}

func main() {
	lunch := Plate{"Fettuccine", "Pumpkin"}

	CookLunch(lunch, SimpleMenu)
	CookDinner(lunch, SimpleMenu)
}
