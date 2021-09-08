package main

import (
	"testing"
)

func Test_getSubject(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "default subject",
			want: "devstaff",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getSubject(); got != tt.want {
				t.Errorf("getSubject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getGreeting(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "default greeting",
			want: "hello",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getGreeting(); got != tt.want {
				t.Errorf("getGreeting() = %v, want %v", got, tt.want)
			}
		})
	}
}
