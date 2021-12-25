package structsMethodsInterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{
		Height: 10.0,
		Width:  10.0,
	}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f , hasArea %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	//checkArea := func(t testing.TB, shape Shape, hasArea float64) {
	//	t.Helper()
	//	got := shape.Area()
	//	if got != hasArea {
	//		t.Errorf("got %g hasArea %g", got, hasArea)
	//	}
	//}
	//t.Run("rectangle", func(t *testing.T) {
	//	rectangle := Rectangle{
	//		Height: 12.0,
	//		Width:  6.0,
	//	}
	//	checkArea(t, rectangle, 72.0)
	//})
	//
	//t.Run("circles", func(t *testing.T) {
	//	circle := Circle{10}
	//	checkArea(t, circle, 314.1592653589793)
	//})
	areaTest := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 12, Width: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTest {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g hasArea %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
