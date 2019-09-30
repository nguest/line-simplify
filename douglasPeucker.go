package main

import (
	"fmt"
	"math"
)

type Datum struct {
	Lon float64
	Lat float64
}

type Line struct {
	V1 Datum
	V2 Datum
}

func diff(v1 Datum, v2 Datum) Datum {
	var vR Datum
	vR.Lon = v2.Lon - v1.Lon
	vR.Lat = v2.Lat - v1.Lat
	return vR
}

func Abs(v1 Datum, v2 Datum) float64 {
	return math.Sqrt(math.Pow(diff(v1, v2).Lon, 2) + math.Pow(diff(v1, v2).Lat, 2))
}

func PerpendicularDistance(v Datum, line Line) float64 {
	x := (line.V2.Lat-line.V1.Lat)*v.Lon - (line.V2.Lon-line.V1.Lon)*v.Lat - line.V2.Lat*line.V1.Lon
	l := Abs(line.V1, line.V2)
	if l == 0 {
		return 0
	}
	return math.Abs(x) / l
}

func DouglasPeucker(data []Datum, e float64) []Datum {
	// Find the point with the maximum distance
	dMax := 0.0
	idx := 0
	end := len(data) - 1
	fmt.Println("datain", data)
	for i := 0; i <= end; i++ {
		line := Line{
			V1: data[0],
			V2: data[end],
		}
		fmt.Println("input", data[i], line)
		d := PerpendicularDistance(data[i], line)
		fmt.Println("d", d)

		if d > dMax {
			idx = i
			dMax = d
		}
	}

	var Res []Datum
	if dMax > e && idx > 1 {
		fmt.Println("dMax", dMax, data, idx)
		// Recursive call
		recR1 := DouglasPeucker(data[0:idx], e)
		recR2 := DouglasPeucker(data[idx:end], e)
		// Build the result list
		fmt.Println(recR1, Res)
		Res = append(recR1[0:len(recR1)-1], recR2[0:len(recR2)]...)
	} else {
		fmt.Println("x", data)
		Res = []Datum{data[0], data[end-1]}
	}
	fmt.Println("Res", Res)
	return Res
}
