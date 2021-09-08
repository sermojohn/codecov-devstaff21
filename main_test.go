package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_doGreet(t *testing.T) {
	t.Parallel()

	type args struct {
		log logger
	}
	tests := []struct {
		name       string
		args       args
		assertFunc func(*testing.T, *args)
	}{
		{
			name: "successful greeting",
			args: args{
				log: &mockLogger{},
			},
			assertFunc: func(t *testing.T, a *args) {
				assert.True(t, a.log.(*mockLogger).DidLog)
				assert.Equal(t, a.log.(*mockLogger).GotFormat, "%s %s!\n")
				assert.Equal(t, a.log.(*mockLogger).GotArgs, []interface{}{"hello", "devstaff"})
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			printGreeting(tt.args.log)
		})
		tt.assertFunc(t, &tt.args)
	}
}

type mockLogger struct {
	DidLog    bool
	GotFormat string
	GotArgs   []interface{}
}

func (m *mockLogger) Logf(format string, a ...interface{}) {
	m.GotFormat = format
	m.GotArgs = a
	m.DidLog = true
}
