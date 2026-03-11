package clockface_test

import (
	"bytes"
	"encoding/xml"
	"slices"
	"testing"
	"time"

	"example.com/learn/clockface"
)

func TestSVGWriterSecondHand(t *testing.T) {
	cases := []struct {
		time time.Time
		line Line
	}{
		{
			simpleTime(0, 0, 0),
			Line{150, 150, 150, 60},
		},
		{
			simpleTime(0, 0, 30),
			Line{150, 150, 150, 150 + 90},
		},
		{
			simpleTime(0, 0, 45),
			Line{150, 150, 60, 150},
		},
		{
			simpleTime(0, 0, 15),
			Line{150, 150, 240, 150},
		},
	}

	for _, c := range cases {
		t.Run(testName(c.time), func(t *testing.T) {
			var b bytes.Buffer
			clockface.SVGWriter(&b, c.time)
			svg := Svg{}
			xml.Unmarshal(b.Bytes(), &svg)

			if !containsLine(c.line, svg.Line) {
				t.Errorf("Expected to find the second hand line %+v, in the SVG lines %+v", c.line, svg.Line)
			}
		})
	}
}

func containsLine(l Line, ls []Line) bool {
	return slices.Contains(ls, l)
}

type Svg struct {
	XMLName xml.Name `xml:"svg"`
	Xmlns   string   `xml:"xmlns,attr"`
	Width   string   `xml:"width,attr"`
	Height  string   `xml:"height,attr"`
	ViewBox string   `xml:"viewBox,attr"`
	Version string   `xml:"version,attr"`
	Circle  Circle   `xml:"circle"`
	Line    []Line   `xml:"line"`
}

type Circle struct {
	Cx float64 `xml:"cx,attr"`
	Cy float64 `xml:"cy,attr"`
	R  float64 `xml:"r,attr"`
}

type Line struct {
	X1 float64 `xml:"x1,attr"`
	Y1 float64 `xml:"y1,attr"`
	X2 float64 `xml:"x2,attr"`
	Y2 float64 `xml:"y2,attr"`
}

func simpleTime(h, m, s int) time.Time {
	return time.Date(2019, time.October, 22, h, m, s, 0, time.UTC)
}

func testName(t time.Time) string {
	return t.Format("15:04:05")
}
