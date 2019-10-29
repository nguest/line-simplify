package algorithms

import (
	"fmt"
	"line-simplify/tracks"
	"reflect"
	"testing"
	"time"
)

func TestDiff(t *testing.T) {
	tn := time.Now()
	v1 := tracks.Datum{0, 10, 0, tn}
	v2 := tracks.Datum{5, 20, 0, tn}
	t.Run("same vector", testFuncDiff(v1, v1, tracks.Datum{Lat: 0, Lon: 0, Alt: 0, Ts: time.Time{}}))
	t.Run("non-zero vectors", testFuncDiff(v1, v2, tracks.Datum{Lat: 5, Lon: 10, Alt: 0, Ts: time.Time{}}))
}

func testFuncDiff(v1, v2, want tracks.Datum) func(*testing.T) {
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
	v0 := tracks.Datum{0, 0, 0, tn}
	v1 := tracks.Datum{0, 1, 0, tn}
	line0 := Line{v0, v0}
	line1 := Line{tracks.Datum{3, 1, 0, tn}, tracks.Datum{1, 5, 0, tn}}
	line2 := Line{tracks.Datum{0, 0, 0, tn}, tracks.Datum{0, 2, 0, tn}}

	t.Run("zeroes", testFuncPerpDist(v0, line0, 0.0))
	t.Run("positive points", testFuncPerpDist(v1, line1, 2.4))
	t.Run("point on line", testFuncPerpDist(v1, line2, 0.0))
}

func testFuncPerpDist(v tracks.Datum, line Line, want float64) func(t *testing.T) {
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
	tn := time.Now()

	data0 := []tracks.Datum{
		{0.0, 0.0, 0, tn}, {5.0, 6.0, 0, tn}, {11.0, 110.0, 0, tn},
		{11.0, 40.0, 0, tn}, {19.0, 12.0, 0, tn}, {22.0, 5.0, 0, tn},
		{21.0, 8.0, 0, tn}, {19.0, 12.0, 0, tn}, {20.0, 20.0, 0, tn},
		{91.0, 15.0, 0, tn}, {19.0, 12.0, 0, tn}, {22.0, 14.0, 0, tn},
	}
	data1 := []tracks.Datum{{0.0, 0.0, 0, tn}, {11.0, 110.0, 0, tn}, {91.0, 15.0, 0, tn}}
	data2 := []tracks.Datum{}
	t.Run("normal data", testFuncDP(data0, data1))
	t.Run("one point", testFuncDP(data2, data2))
}

func testFuncDP(data []tracks.Datum, want []tracks.Datum) func(*testing.T) {
	return func(t *testing.T) {
		got := DPByTolerance(data, 5.5)
		if !reflect.DeepEqual(got, want) {
			fmt.Printf("WANT: %+v\n GOT: %+v\n", want, got)
			t.FailNow()
		}
	}
}
