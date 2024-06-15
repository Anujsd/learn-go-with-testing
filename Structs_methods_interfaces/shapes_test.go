package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{3.2, 3.4}
	got := Perimeter(rectangle)
	want := 13.2
	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// func TestArea(t *testing.T) {

// 	checkArea := func(t testing.TB, shape Shape, want float64) {
// 		got := shape.Area()
// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	}
// 	t.Run("Reactange", func(t *testing.T) {
// 		rectangle := Rectangle{3.2, 3.4}
// 		want := 10.88
// 		checkArea(t, rectangle, want)
// 	})

// 	t.Run("Circle", func(t *testing.T) {
// 		circle := Circle{3.2}
// 		want := 32.169908772759484
// 		checkArea(t, circle, want)
// 	})
// }

// Table Driven Development
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 3.2, Height: 3.4}, hasArea: 10.88},
		{name: "Circle", shape: Circle{radius: 3.2}, hasArea: 32.169908772759484},
		{name: "Triangle", shape: Triangle{base: 12, height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			want := tt.hasArea
			if got != want {
				t.Errorf("%#v got %g want %g", tt.shape, got, want)
			}
		})

	}
}
