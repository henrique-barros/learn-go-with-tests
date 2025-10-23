package main

import "testing"

func TestPerimeter(t *testing.T) {
	got := Perimeter(Rectangle{10.0, 10.0})
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := map[string]struct {
		shape Shape
		want  float64
	}{
		"circles": {
			shape: Circle{10.0},
			want:  314.1592653589793,
		},
		"rectangles": {
			shape: Rectangle{10.0, 10.0},
			want:  100.0,
		},
		"triangles": {
			shape: Triangle{12, 6},
			want:  36.0,
		},
	}

	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("%#v got %.2f want %.2f", shape, got, want)
		}
	}

	for key, value := range areaTests {
		t.Run(key, func(t *testing.T) {
			t.Parallel()
			checkArea(t, value.shape, value.want)
		})
	}
}
