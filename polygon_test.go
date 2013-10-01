package geofence

import (
	"testing"
)

func TestPolygonCopy(t *testing.T) {
	startPolygon := Polygon{Point{1, 2}, Point{3, 4}, Point{5, 6}}
	cpPolygon := startPolygon.copy()
	for i, pt := range startPolygon {
		if cpPolygon[i] != pt {
			t.Error("copy failed")
		}
	}
}
