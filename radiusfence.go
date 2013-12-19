package geofence

import (
	"math"
)

const (
	DegToRad = 0.017453292519943295769236907684886127134428718885417 // N[Pi/180, 50]
	RadToDeg = 57.295779513082320876798154814105170332405472466564   // N[180/Pi, 50]

	GradToRad = 0.015707963267948966192313216916397514420985846996876 // N[Pi/200, 50]
	RadToGrad = 63.661977236758134307553505349005744813783858296183   // N[200/Pi, 50]

	Radius = 6371 //Earth in km

	North = 0
	East  = 90
	South = 180
	West  = 270
)

type Radiusfence struct {
	p Point
	r float64
}

// NewRadiusfence returns a new Radiusfence for the given slice of points
func NewRadiusfence(p Point, r float64) *Radiusfence {
	rf := &Radiusfence{
		p: p,
		r: r,
	}

	return rf
}

// Inside returns whether the given point lies inside the Radiusfence
func (rf *Radiusfence) Inside(pt Point) bool {
	//Given point falls inside the radius
	if calculateInside(*rf, pt) {
		return true
	}

	//Not inside
	return false
}

func calculateInside(rf Radiusfence, pt Point) bool {
	//Calculate bounding box

	x1, x2, y1, y2 := getBoundingBox(rf)
	if rf.p.Latitude > x1 && rf.p.Latitude < x2 {
		if rf.p.Longitude > y1 && rf.p.Longitude < y2 {
			//Inside bounding box
			return pointInCircle(rf, pt)
		}
	}

	return false
}

func getBoundingBox(rf Radiusfence) (x1, x2, y1, y2 float64) {
	var lat1, lat2, lon1, lon2 float64

	//Convert long,lat to rad
	latRad := rf.p.Latitude * DegToRad
	longRad := rf.p.Longitude * DegToRad

	northMost := math.Asin(math.Sin(latRad)*math.Cos(rf.r/Radius) + math.Cos(latRad)*math.Sin(rf.r/Radius)*math.Cos(North))
	southMost := math.Asin(math.Sin(latRad)*math.Cos(rf.r/Radius) + math.Cos(latRad)*math.Sin(rf.r/Radius)*math.Cos(South))
	eastMost := longRad + math.Atan2(math.Sin(East)*math.Sin(rf.r/Radius)*math.Cos(latRad), math.Cos(rf.r/Radius)-math.Sin(latRad)*math.Sin(latRad))
	westMost := longRad + math.Atan2(math.Sin(West)*math.Sin(rf.r/Radius)*math.Cos(latRad), math.Cos(rf.r/Radius)-math.Sin(latRad)*math.Sin(latRad))

	if northMost > southMost {
		lat1 = southMost
		lat2 = northMost
	} else {
		lat1 = northMost
		lat2 = southMost
	}

	if eastMost > westMost {
		lon1 = westMost
		lon2 = eastMost
	} else {
		lon1 = eastMost
		lon2 = westMost
	}

	return lat1, lat2, lon1, lon2
}

func pointInCircle(rf Radiusfence, pt Point) bool {
	distance := ((math.Acos(
		(math.Sin(rf.p.Latitude*math.Pi/180)*math.Sin(pt.Latitude*math.Pi)/180)+
			(math.Cos(rf.p.Latitude*math.Pi/180)*math.Cos(pt.Latitude*math.Pi)/180)*
				(math.Cos((rf.p.Longitude-pt.Latitude)*math.Pi/180))) * 180 / math.Pi) * 60 * 1.1515)

	if distance < rf.r {
		return true
	}

	return false
}
