package main

import (
	"github.com/sermojohn/codecov-devstaff21/greet"
	"github.com/sermojohn/codecov-devstaff21/log"
)

func main() {
	logger := log.NewFmtLogger()
	doGreet(logger)
}

func doGreet(logger Logger) {
	var (
		greeter  = &greet.Greeter{}
		greeting = greeter.GetGreeting()
		subject  = greeter.GetSubject()
	)
	logger.Printf("%s %s!\n", greeting, subject)
}

type Logger interface {
	Printf(format string, a ...interface{}) (n int, err error)
}
