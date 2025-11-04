package struct_method_interface

import "testing"

func Test_perimeter(t *testing.T) {
	rectangle := Rectangle{2, 4}
	got := rectangle.Perimeter()
	want := 12.0

	assertCorrectMessage(t, got, want)
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle
		want := 72.0

		checkArea(t, got, want)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle
		want := 314.1592653589793

		checkArea(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
