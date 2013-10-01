package geofence

import (
	"testing"
)

func TestPolyfence(t *testing.T) {
	thePolyfence := NewPolyfence(Polygon{
		{0, 6},
		{2, -2},
		{4, 6},
		{6, -2},
		{8, 6},
		{12, 0},
		{14, 8},
		{10, 10},
		{6, 4},
		{4, 10},
	})

	testPointsInside := []Point{
		{2, -1},
		{2, 7},
		{4, 8},
		{5, 4.5},
		{6, 1},
	}

	for _, pt := range testPointsInside {
		if !thePolyfence.Inside(pt) {
			t.Errorf("Expected point %v to be inside poly. Got outside.", pt)
		}
	}

	testPointsOutside := []Point{
		// ridicilous values
		{-100, -100},
		{-100, 100},
		{100, -100},
		{100, 100},

		// near values
		{0, 0},
		{2, 9},
		{3, 1},
		{4, 5},
		{5, 8},
		{6, 5},
		{6, 9},
		{6, 10},
		{8, -1},
		{8, 0},
		{8, 5},
		{8, 9},
		{13, 3.5},
		{13, 9},
		{14, 0},
		{14, 10},
		{15, 0},
		{15, 8},
		{15, 10},
	}

	for _, pt := range testPointsOutside {
		if thePolyfence.Inside(pt) {
			t.Errorf("Expected point %v to be outisde poly. Got inside.", pt)
		}
	}
}
