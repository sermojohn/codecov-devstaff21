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

	printGreeting(log, true)
}

func printGreeting(l logger, isMorning bool) {
	var (
		greeter  = &greet.Greeter{}
		subject  = greeter.GetSubject()
		greeting = greeter.GetGeneralGreeting()
	)

	if isMorning {
		greeting = greeter.GetMorningGreeting()
	}
	l.Logf("%s, %s!\n", greeting, subject)
}

type logger interface {
	Logf(format string, a ...interface{})
}

type factory struct{}

func (f *factory) NewLogger() logger {
	return &log.FmtLogger{PrintfFunc: fmt.Printf}
}
