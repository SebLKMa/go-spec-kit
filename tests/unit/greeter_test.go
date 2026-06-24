package unit_test

import (
	"testing"

	"github.com/sebmaspd/hello-spec-it/internal/greeter"
)

func TestGreet_SingleName(t *testing.T) {
	got := greeter.Greet("World")
	want := "hello World"
	if got != want {
		t.Errorf("Greet(\"World\") = %q, want %q", got, want)
	}
}

func TestGreet_MultiWordName(t *testing.T) {
	got := greeter.Greet("Jane Doe")
	want := "hello Jane Doe"
	if got != want {
		t.Errorf("Greet(\"Jane Doe\") = %q, want %q", got, want)
	}
}

func TestGreet_SpecialCharacters(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{"O'Brien", "hello O'Brien"},
		{"José", "hello José"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := greeter.Greet(tt.name)
			if got != tt.want {
				t.Errorf("Greet(%q) = %q, want %q", tt.name, got, tt.want)
			}
		})
	}
}
