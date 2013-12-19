package geofence

import (
	"testing"
)

func testRadius(t *testing.T) {
	radius := 350.0
	insidePoint := &Point{
		52.628218,
		4.754505,
	}

	outsidePoint := &Point{
		52.628674,
		4.760942,
	}

	radiusFence := NewRadiusfence(*insidePoint, radius)
	if calculateInside(*radiusFence, *insidePoint) == false {
		t.Error("Point is not inside radius")
	}

	if calculateInside(*radiusFence, *outsidePoint) == true {
		t.Error("Point is not outside radius")
	}

	t.Log("Succesfully passed the tests")
}
