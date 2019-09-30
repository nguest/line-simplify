package main

import (
	"fmt"
	"testing"
)

func TestPerpendicularDistance(t *testing.T) {
	v := Datum{2, 1}
	line := Line{Datum{0, 0}, Datum{4, 0}}
	got := PerpendicularDistance(v, line)
	want := 1.0
	if got != want {
		fmt.Printf("WANT: %+v\n GOT: %+v\n", want, got)
		t.FailNow()
	}
}

func TestAbs(t *testing.T) {
	v1 := Datum{0, 2}
	v2 := Datum{3, 6}
	got := Abs(v1, v2)
	want := 5.0
	if got != want {
		fmt.Printf("WANT: %+v\n GOT: %+v\n", want, got)
		t.FailNow()
	}
}
