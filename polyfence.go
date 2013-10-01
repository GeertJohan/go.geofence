package geofence

// The Polyfence is based on information and explanation by softSurfer (Dan Sunday)
// http://geomalgorithms.com/a03-_inclusion.html

// Polyfence is a geofence defined by a set of points (Polygon).
type Polyfence struct {
	p Polygon
}

// NewPolyfence returns a new Polyfence for the given slice of points
func NewPolyfence(p Polygon) *Polyfence {
	// create new polyfence with copy of given vertex
	pf := &Polyfence{
		p: p.copy(),
	}

	// complete the vertex (if necicary)
	if len(pf.p) > 0 && pf.p[0] != pf.p[len(pf.p)-1] {
		pf.p = append(pf.p, pf.p[0])
	}

	// all done
	return pf
}

// Inside returns whether the given point lies inside the Polyfence
// The used algorithm is Winding Number
func (pf *Polyfence) Inside(pt Point) bool {
	// given point lies outside the polygon when winding number equals 0
	if calculateWindingNumber(pf.p, pt) == 0 {
		return false
	}

	// all other values: inside
	return true
}

// perform a winding number calculation on given vertex (polygon) and point
func calculateWindingNumber(p Polygon, pt Point) int {
	// quick return for non-poly vertex
	if len(p) < 3 {
		return 0
	}

	// the winding number counter
	wn := 0

	// amount of edges to check
	n := len(p) - 2

	// loop through all edges of the polygon
	for i := 0; i <= n; i++ {
		// start y <= pt.Latitude
		if p[i].Latitude <= pt.Latitude {
			// an upward crossing
			if p[i+1].Latitude > pt.Latitude {
				// P left of  edge
				if isLeft(p[i], p[i+1], pt) > 0 {
					// have  a valid up intersect
					wn++
				}
			}
		} else { // p[i].Latitude > pt.Latitude
			// a downward crossing
			if p[i+1].Latitude <= pt.Latitude {
				// P right of edge
				if isLeft(p[i], p[i+1], pt) < 0 {
					// have  a valid down intersect
					wn--
				}
			}
		}
	}

	// all done
	return wn
}

func isLeft(lineA Point, lineB Point, pt Point) float64 {
	return ((lineB.Longitude-lineA.Longitude)*(pt.Latitude-lineA.Latitude) - (pt.Longitude-lineA.Longitude)*(lineB.Latitude-lineA.Latitude))
}
