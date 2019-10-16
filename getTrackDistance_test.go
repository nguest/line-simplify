package main

import (
	"fmt"
	"testing"
)

func TestGetTotalTrackLength(t *testing.T) {
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
	t.Run("zeroes", testFuncGetTotalTrackLength(data0, 0))
	t.Run("positive points", testFuncGetTotalTrackLength(data1, 850.8496136907015))
	t.Run("negative points", testFuncGetTotalTrackLength(data2, 850.8496136907015))
}

func testFuncGetTotalTrackLength(data []Datum, want float64) func(*testing.T) {
	return func(t *testing.T) {
		got := GetTotalTrackLength(data)
		if got != want {
			fmt.Printf("WANT: %+v\n GOT: %+v\n", want, got)
			t.FailNow()
		}
	}
}

// func testSumFunc(numbers []int, expected int) func(*testing.T) {
// 	return func(t *testing.T) {
// 		actual := Sum(numbers)
// 		if actual != expected {
// 			t.Error(fmt.Sprintf("Expected the sum of %v to be %d but instead got %d!", numbers, expected, actual))
// 		}
// 	}
