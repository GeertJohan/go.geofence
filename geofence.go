package geofence

// Geofence abstracts any type of geofence through an interface.
type Geofence interface {
	Inside(p Point) bool
}

//++ http://geomalgorithms.com/a03-_inclusion.html
