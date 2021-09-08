package greet

import (
	"testing"
)

func TestGreeter_GetGeneralGreeting(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		g    *Greeter
		want string
	}{
		{
			name: "default greeting",
			g:    &Greeter{},
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.GetGeneralGreeting(); got != tt.want {
				t.Errorf("Greeter.GetGeneralGreeting() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreeter_GetSubject(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		g    *Greeter
		want string
	}{
		{
			name: "default subject",
			g:    &Greeter{},
			want: "devstaff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.GetSubject(); got != tt.want {
				t.Errorf("Greeter.GetSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGreeter_GetMorningGreeting(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		g    *Greeter
		want string
	}{
		{
			name: "morning greeting",
			g:    &Greeter{},
			want: "good morning",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.GetMorningGreeting(); got != tt.want {
				t.Errorf("Greeter.GetMorningGreeting() = %v, want %v", got, tt.want)
			}
		})
	}
}
