package clockface_test

import (
	"testing"
	"time"

	"example.com/learn/clockface"
)

func TestSecondHandAtMidnight(t *testing.T) {
	time := time.Date(2019, time.October, 22, 0, 0, 0, 0, time.UTC)
	want := clockface.Point{150, 150 - 90}
	got := clockface.SecondHand(time)

	if want != got {
		t.Errorf("Got %v, Want %v", got, want)
	}
}

func TestSecondHandAt30Seconds(t *testing.T) {
	time := time.Date(2019, time.October, 22, 0, 0, 30, 0, time.UTC)
	want := clockface.Point{150, 150 + 90}
	got := clockface.SecondHand(time)

	if want != got {
		t.Errorf("Got %v, Want %v", got, want)
	}
}
