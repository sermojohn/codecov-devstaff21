package main

import (
	"fmt"

	"github.com/sermojohn/codecov-devstaff21/greet"
)

func main() {
	doGreet()
}

func doGreet() {
	greeter := greet.Greeter{}
	fmt.Printf("%s %s!\n", greeter.GetGreeting(), greeter.GetSubject())
}
