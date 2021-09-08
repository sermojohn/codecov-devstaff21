package log

import "testing"

func Test_fmtLogger_Printf(t *testing.T) {
	type fields struct {
		printfFunc func(format string, a ...interface{}) (n int, err error)
	}
	type args struct {
		format string
		a      []interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantN   int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &fmtLogger{
				printfFunc: tt.fields.printfFunc,
			}
			gotN, err := f.Printf(tt.args.format, tt.args.a...)
			if (err != nil) != tt.wantErr {
				t.Errorf("fmtLogger.Printf() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotN != tt.wantN {
				t.Errorf("fmtLogger.Printf() = %v, want %v", gotN, tt.wantN)
			}
		})
	}
}
