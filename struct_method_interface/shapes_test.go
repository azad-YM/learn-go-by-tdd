package struct_method_interface

import "testing"

func Test_perimeter(t *testing.T) {
	rectangle := Rectangle{2, 4}
	got := rectangle.Perimeter()
	want := 12.0

	assertCorrectMessage(t, got, want)
}

type TableTests struct {
	name    string
	shape   Shape
	hasArea float64
}

func TestArea(t *testing.T) {
	areaTests := []TableTests{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})

	}
}

func assertCorrectMessage(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %g want %g", got, want)
	}
}
