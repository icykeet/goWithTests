package structure

import (
	"testing"
)

// func TestPerimetr(t *testing.T) {
// 	t.Run("test Rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{10.0, 10.0}
// 		got := rectangle.Perimeter()
// 		want := 40.0
// 		if got != want {
// 			t.Errorf("got %.2f, want %.2f", got, want)
// 		}
// 	})
// }

// func TestArea(t *testing.T) {
// 	checkArea := func(t *testing.T, shape Shape, want float64) {
// 		t.Helper()
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %.2g, but want %.2g", got, want)
// 		}
// 	}
// 	t.Run("test Rectangle", func(t *testing.T) {
// 		rectangle := Rectangle{12.0, 6.0}
// 		checkArea(t, rectangle, 72.0)
// 	})
// 	t.Run("test Cirlce", func(t *testing.T) {
// 		circle := Circle{10}
// 		checkArea(t, circle, 314.1592653589793)
// 	})
// }

func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// using tt.name from the case to use it as the `t.Run` test name
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}

func TestPerimeter(t *testing.T) {
	perimeterTest := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 36},
		{Circle{10}, 62.83185307179586},
	}

	for _, tt := range perimeterTest {
		got := tt.shape.Perimeter()
		if got != tt.want {
			t.Errorf("got %g, but want %g", got, tt.want)
		}
	}
}

func BenchmarkTestPeimeter(b *testing.B) {
	rectangle := Rectangle{12, 6}
	for i := 0; i < b.N; i++ {
		rectangle.Area()
	}
}
