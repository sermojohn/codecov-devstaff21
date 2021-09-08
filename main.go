package main

import "fmt"

func main() {
	greet()
}

func greet() {
	fmt.Printf("%s %s!\n", getGreeting(), getSubject())
}

func getGreeting() string {
	return "hello"
}

func getSubject() string {
	return "devstaff"
}
