package dependency_injection

import (
	"testing"
)

func TestGreet(t *testing.T) {
	log := Log{message: ""}
	Greet(&log, "Azad")

	got := log.message
	want := "Hello, Azad"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
