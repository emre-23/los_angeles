package messagesmes_test

import (
	"testing"

	"github.com/emre-23/los_angeles/messages"
)

func TestGreet(t *testing.T) {
	got := messages.Greet("Gopher")
	expected := "Hello, Gopher!\n"
	if got != expected {
		t.Errorf("Yanlış. Got %v, want %v", got, expected)
	}
}

func TestDepart(t *testing.T) {
	got := messages.Depart("Gopher")
	expected := "Goodbye, Gopher\n"
	if got != expected {
		t.Errorf("Yanlış. Got %v, want %v", got, expected)
	}
}
