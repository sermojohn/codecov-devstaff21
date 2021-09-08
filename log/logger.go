package log

import "fmt"

func NewFmtLogger() *fmtLogger {
	return &fmtLogger{printfFunc: fmt.Printf}
}

type fmtLogger struct {
	printfFunc func(format string, a ...interface{}) (n int, err error)
}

func (f *fmtLogger) Printf(format string, a ...interface{}) (n int, err error) {
	return f.printfFunc(format, a...)
}
