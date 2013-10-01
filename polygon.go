package geofence

// Polygon defines a slice of Point's that indicate an array
type Polygon []Point

// copy returns a copy of the vertex
func (p Polygon) copy() Polygon {
	// create a new vertex with the cap to copy existing vertex into it
	//++ TODO: should newPolygon have cap=len(v) or cap=cap(v) ?
	newPolygon := make(Polygon, len(p))
	// use builtin copy
	copy(newPolygon, p)
	// return new vertex
	return newPolygon
}
