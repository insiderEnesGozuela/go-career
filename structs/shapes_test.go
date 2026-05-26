package structs

import (
	"testing"
)

type TestCase struct {
	shape Shape
	want  float64
}

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 10.0}

		got := rectangle.Perimeter()
		want := 40.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{5.0}

		got := circle.Perimeter()
		want := 2 * 3 * 5.0

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Rect", shape: Rectangle{Width: 12, Height: 6}, want: 72.0},
		{name: "Circ", shape: Circle{Radius: 10}, want: 300.0},
		{name: "Tria", shape: Triangle{Height: 10.0, Base: 2.0}, want: 10.0},
	}

	for _, areaTest := range areaTests {
		t.Run(areaTest.name, func(t *testing.T) {
			got := areaTest.shape.Area()

			if got != areaTest.want {
				t.Errorf("got %.2f want %.2f", got, areaTest.want)
			}
		})
	}
}
