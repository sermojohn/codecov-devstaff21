package main

import (
	"fmt"

	"github.com/sermojohn/codecov-devstaff21/greet"
	"github.com/sermojohn/codecov-devstaff21/log"
)

func main() {
	var (
		f   = &factory{}
		log = f.NewLogger()
	)

	printGreeting(log)
}

func printGreeting(l logger) {
	var (
		greeter  = &greet.Greeter{}
		greeting = greeter.GetGreeting()
		subject  = greeter.GetSubject()
	)
	l.Logf("%s %s!\n", greeting, subject)
}

type logger interface {
	Logf(format string, a ...interface{})
}

type factory struct{}

func (f *factory) NewLogger() logger {
	return &log.FmtLogger{PrintfFunc: fmt.Printf}
}
