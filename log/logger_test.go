package log

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_fmtLogger_Logf(t *testing.T) {
	t.Parallel()

	type fields struct {
		fmtPrinter *mockFmtPrinter
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name       string
		fields     fields
		args       args
		wantN      int
		wantErr    error
		assertFunc func(*testing.T, *fields)
	}{
		{
			name: "successful printf",
			args: args{
				format: "%s %s",
				a:      []interface{}{"hello", "world"},
			},
			fields: fields{
				fmtPrinter: &mockFmtPrinter{},
			},
			assertFunc: func(t *testing.T, f *fields) {
				assert.True(t, f.fmtPrinter.DidPrint)
			},
		},
		{
			name: "failed printf",
			args: args{
				format: "%s %s",
				a:      []interface{}{"hello", "world"},
			},
			fields: fields{
				fmtPrinter: &mockFmtPrinter{PrintError: errors.New("printf error")},
			},
			wantErr: errors.New("printf error"),
			assertFunc: func(t *testing.T, f *fields) {
				assert.False(t, f.fmtPrinter.DidPrint)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FmtLogger{
				PrintfFunc: tt.fields.fmtPrinter.Printf,
			}
			f.Logf(tt.args.format, tt.args.a...)

			if tt.assertFunc != nil {
				tt.assertFunc(t, &tt.fields)
			}
		})
	}
}

type mockFmtPrinter struct {
	DidPrint   bool
	PrintError error
}

func (m *mockFmtPrinter) Printf(format string, a ...interface{}) (n int, err error) {
	if m.PrintError != nil {
		return 0, m.PrintError
	}
	m.DidPrint = true
	return 0, nil
}
