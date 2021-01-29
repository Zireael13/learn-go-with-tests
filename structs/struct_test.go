package structs

import "testing"

func assert_eq_floats(t testing.TB, got, want float64) {
	t.Helper()
	if got != want {
		t.Errorf("got %.2f want %.2F", got, want)
	}
}

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasPerimeter: 36},
		{name: "Circle", shape: Circle{10}, hasPerimeter: 62.83185307179586476925},
		{name: "Triangle", shape: Triangle{8, 6, 4}, hasPerimeter: 18.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasPerimeter)
			}

		})
	}

	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	assert_eq_floats(t, got, want)

}

func TestArea(t *testing.T) {
	// checkArea := func(t testing.TB, shape Shape, want float64) {
	// 	t.Helper()
	// 	got := shape.Area()
	// 	assert_eq_floats(t, got, want)
	// }

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{24, 30, 18}, hasArea: 216.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

	// t.Run("rectangle", func(t *testing.T) {
	// 	rect := Rectangle{10.0, 10.0}
	// 	want := 100.0

	// 	checkArea(t, rect, want)
	// })

	// t.Run("circle", func(t *testing.T) {
	// 	circle := Circle{10}
	// 	want := 314.1592653589793
	// 	checkArea(t, circle, want)
	// })
}
