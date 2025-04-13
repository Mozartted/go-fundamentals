package structsmethodsinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	rctangle := Rectangle{10.0, 10.0}
	got := rctangle.Perimeter()

	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()

		if want != got {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	}

	t.Run("Test rectangle area", func(t *testing.T) {
		rctangle := Rectangle{12.0, 6.0}
		want := 72.0
		checkArea(t, rctangle, want)
	})
}
