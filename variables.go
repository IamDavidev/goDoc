package main

import "fmt"

// variable de scope global
var i = 100

// variable de scope para todo el programa || proyecto
var J = 200

func main() {

	var (
		david  string = "davide"
		lezama int    = 10
		trejo  bool   = true
	)

	var (
		user   string = "itsDavidev"
		emial  string = "davidlez2mgk@gmail.com"
		active bool   = true
	)

	if active {
		fmt.Println("user => ", user)
		fmt.Println("email => ", emial)
		fmt.Println("active => ", active)

	}

	fmt.Println(david, lezama, trejo)

}
