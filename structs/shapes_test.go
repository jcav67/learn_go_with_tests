package shapes

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{Width: 10.0, Height: 10.0}
	got := rectangle.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
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
		checkArea(t, rectangle, 72.0)
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}

func TestAreaTableTest(t *testing.T) {

	areaTests := []struct {
		testName     string
		shape        Shape
		expectedArea float64
	}{
		{shape: Rectangle{Width: 12, Height: 6}, expectedArea: 73.0},
		{shape: Circle{Radius: 10}, expectedArea: 314.1592653589793},
		{shape: Triangle{Base: 12, Height: 6}, expectedArea: 36.0},
	}

	for index, tt := range areaTests {
		t.Run(tt.testName, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.expectedArea {
				index += 1
				t.Errorf("%#v got %g want %g, error found in test number %v", tt.shape, got, tt.expectedArea, index)
			}
		})
	}

}

// --------------------------------- forma más larga de hacer tests ------------------------

// func TestArea(t *testing.T) {
// 	t.Run("rectangles", func(t *testing.T) {
// 		rectangle := Rectangle{12, 6}
// 		got := rectangle.Area()
// 		want := 72.0

// 		if got != want {
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})

// 	t.Run("circles", func(t *testing.T) {
// 		circle := Circle{10}
// 		got := circle.Area()
// 		want := 314.1592653589793

// 		if got != want {
// 			// %g da mas decimales
// 			t.Errorf("got %g want %g", got, want)
// 		}
// 	})
// }
