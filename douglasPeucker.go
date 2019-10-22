package main

import (
	"math"
	"sort"
	"time"
)

type Datum struct {
	Lat float64   `json:"lat"`
	Lon float64   `json:"lon"`
	Alt float64   `json:"alt"`
	Ts  time.Time `json:"ts"`
}

type Line struct {
	V1 Datum
	V2 Datum
}

func Diff(v1 Datum, v2 Datum) Datum {
	var vR Datum
	vR.Lon = v2.Lon - v1.Lon
	vR.Lat = v2.Lat - v1.Lat
	return vR
}

func Abs(v1 Datum, v2 Datum) float64 {
	return math.Sqrt(math.Pow(Diff(v1, v2).Lon, 2) + math.Pow(Diff(v1, v2).Lat, 2))
}

func PerpendicularDistance(v Datum, line Line) float64 {
	if v == line.V1 || v == line.V2 || line.V1 == line.V2 {
		return 0.0
	}
	x := (line.V2.Lat-line.V1.Lat)*v.Lon - (line.V2.Lon-line.V1.Lon)*v.Lat + line.V2.Lon*line.V1.Lat - line.V2.Lat*line.V1.Lon
	l := Abs(line.V2, line.V1)
	if l == 0 {
		return 0
	}
	return math.Abs(x) / l
}

// DouglasPeucker simple Douglas-Peucker algorithm with tolerance e
func DouglasPeucker(data []Datum, e float64) []Datum {
	// Find the point with the maximum distance
	dMax := 0.0
	idx := 0
	end := len(data) - 1

	for i := 0; i <= end; i++ {
		line := Line{
			V1: data[0],
			V2: data[end],
		}
		d := PerpendicularDistance(data[i], line)

		if d > dMax {
			idx = i
			dMax = d
		}
	}

	var Res []Datum

	if dMax > e && idx > 1 {
		// Recursive call
		recR1 := DouglasPeucker(data[0:idx], e)
		recR2 := DouglasPeucker(data[idx:end], e)
		// Build the result list
		Res = append(recR1[0:len(recR1)-1], recR2[0:len(recR2)]...)
	} else {
		if len(data) > 1 {
			Res = []Datum{data[0], data[end-1]}
		} else {
			Res = data
		}
	}
	return Res
}

// DPByCount implements Douglas Peucker but with  a given pointscount for the result set
func DPByCount(data []Datum, count int) []Datum {
	len := len(data)
	weights := make([]float64, len)
	var dP func(int, int)

	dP = func(start, end int) {
		if end <= start+1 {
			return
		}

		line := Line{
			V1: data[start],
			V2: data[end],
		}
		dMax := -1.0
		idx := 0
		d := 0.0

		for i := start + 1; i < end; i++ {
			d = PerpendicularDistance(data[i], line)
			if d > dMax {
				dMax = d
				idx = i
			}
		}

		weights[idx] = dMax

		dP(start, idx)
		dP(idx, end)
	}

	dP(0, len-1)

	// make sure first and last point always included
	weights[0] = math.MaxFloat64
	weights[len-1] = math.MaxFloat64

	// sort []weights descending, to calculate maxT max tolerance
	weightsDesc := make([]float64, len)
	copy(weightsDesc, weights)
	sort.Slice(weightsDesc, func(i, j int) bool {
		return weightsDesc[i] > weightsDesc[j]
	})
	maxT := weightsDesc[count-1]

	// filter correct highest-weighted points into []dataOut
	n := 0
	dataOut := make([]Datum, count)
	for i, x := range data {
		if weights[i] >= maxT {
			dataOut[n] = x
			n++
		}
	}

	return dataOut
}
