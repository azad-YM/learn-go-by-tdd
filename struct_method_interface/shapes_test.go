package struct_method_interface

import "testing"

func Test_perimeter(t *testing.T) {
	rectangle := Rectangle{2, 4}
	got := rectangle.Perimeter()
	want := 12.0

	assertCorrectMessage(t, got, want)
}

type TableTests struct {
	shape Shape
	want  float64
}

func TestArea(t *testing.T) {
	testsArea := []TableTests{
		{shape: Rectangle{12, 6}, want: 72},
		{shape: Circle{10}, want: 314.1592653589793},
	}

	for _, tt := range testsArea {
		got := tt.shape.Area()
		want := tt.want

		assertCorrectMessage(t, got, want)
	}
}

func assertCorrectMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
