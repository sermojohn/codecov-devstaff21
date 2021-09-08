package log

type FmtLogger struct {
	PrintfFunc func(format string, a ...interface{}) (n int, err error)
}

func (f *FmtLogger) Logf(format string, a ...interface{}) {
	f.PrintfFunc(format, a...)
}
