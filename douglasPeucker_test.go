package main

import (
	"fmt"
	"testing"
	"time"
)

func TestDiff(t *testing.T) {
	tn := time.Now()
	v1 := Datum{0, 10, 0, tn}
	v2 := Datum{5, 20, 0, tn}
	t.Run("same vector", testFuncDiff(v1, v1, Datum{Lat: 0, Lon: 0, Alt: 0, Ts: time.Time{}}))
	t.Run("non-zero vectors", testFuncDiff(v1, v2, Datum{Lat: 5, Lon: 10, Alt: 0, Ts: time.Time{}}))
}

func testFuncDiff(v1, v2, want Datum) func(*testing.T) {
	return func(t *testing.T) {
		got := Diff(v1, v2)
		if got != want {
			fmt.Printf("WANT: %+v\n GOT: %+v\n", want, got)
			t.FailNow()
		}
	}
}

func TestPerpendicularDistance(t *testing.T) {
	tn := time.Now()
	v0 := Datum{0, 0, 0, tn}
	v1 := Datum{0, 1, 0, tn}
	line0 := Line{v0, v0}
	line1 := Line{Datum{3, 1, 0, tn}, Datum{1, 5, 0, tn}}
	line2 := Line{Datum{0, 0, 0, tn}, Datum{0, 2, 0, tn}}

	t.Run("zeroes", testFuncPerpDist(v0, line0, 0.0))
	t.Run("positive points", testFuncPerpDist(v1, line1, 2.4))
	t.Run("point on line", testFuncPerpDist(v1, line2, 0.0))
}

func testFuncPerpDist(v Datum, line Line, want float64) func(t *testing.T) {
	return func(t *testing.T) {
		got := PerpendicularDistance(v, line)
		if got != want {
			fmt.Printf("WANT: %+v\n GOT: %+v\n", want, got)
			t.FailNow()
		}
	}
}

func TestDouglasPeucker(t *testing.T) {
	//t.Run
	data0 := []Datum{}
	data1 := []Datum{
		Datum{
			Lat: 21.0,
			Lon: 10.0,
		},
		Datum{
			Lat: 18.0,
			Lon: 12.0,
		},
		Datum{
			Lat: 22.0,
			Lon: 11.0,
		},
	}
	data2 := []Datum{
		Datum{
			Lat: -21.0,
			Lon: -10.0,
		},
		Datum{
			Lat: -18.0,
			Lon: -12.0,
		},
		Datum{
			Lat: -22.0,
			Lon: -11.0,
		},
	}

	t.Run("zeroes", testFuncDP(data0, 0))
	t.Run("positive points", testFuncDP(data1, 850.8496136907015))
	t.Run("negative points", testFuncDP(data2, 850.8496136907015))
}

func testFuncDP(data []Datum, want float64) func(*testing.T) {
	return func(t *testing.T) {
		got := GetTotalTrackLength(data)
		if got != want {
			fmt.Printf("WANT: %+v\n GOT: %+v\n", want, got)
			t.FailNow()
		}
	}
}
