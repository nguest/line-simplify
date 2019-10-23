package tracks

import (
	"math"
)

// GetTotalTrackLength is the sum of the length of all track segments
func GetTotalTrackLength(data []*Datum) float64 {
	totalLen := 0.0
	for i := 0; i < len(data)-1; i++ {
		segmentLen := haversine(data[i], data[i+1])
		totalLen += segmentLen
	}
	return totalLen
}

func degToRad(deg float64) float64 {
	return math.Pi * deg / 180
}

// haversine computes the distance between two points in km taking into account the curve of the earth
func haversine(p1, p2 *Datum) float64 {
	R := 6371.0 // km
	φ1 := degToRad(p1.Lat)
	φ2 := degToRad(p2.Lat)
	Δφ := degToRad(p2.Lat - p1.Lat)
	Δλ := degToRad(p2.Lon - p1.Lon)

	a := math.Sin(Δφ/2)*math.Sin(Δφ/2) + math.Cos(φ1)*math.Cos(φ2)*math.Sin(Δλ/2)*math.Sin(Δλ/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))

	return R * c
}
