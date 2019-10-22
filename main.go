package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	in, err := ReadAndParse("sampleTracks/sample1.igc")

	outDP := DPByCount(in.Data, 5)
	outVis := Visvalingam(in.Data, 5)

	var L struct {
		Line [][]float64 `json:"line"`
		Dist float64     `json:"dist"`
	}

	var X struct {
		In      [][]float64 `json:"in"`
		OutDP   [][]float64 `json:"outDP"`
		OutVis  [][]float64 `json:"outVis"`
		OutDist float64     `json:"outDist"`
	}
	for _, v := range outDP {
		X.OutDP = append(X.OutDP, []float64{v.Lon, v.Lat})
	}
	for _, v := range outVis {
		X.OutVis = append(X.OutVis, []float64{v.Lon, v.Lat})
	}
	for _, v := range in.Data {
		X.In = append(X.In, []float64{v.Lon, v.Lat})
	}
	X.OutDist = GetTotalTrackLength(outDP)
	json.Marshal(X)
	t.Execute(w, X)
}

func handleRequests() {
	r := mux.NewRouter().StrictSlash(true)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	r.HandleFunc("/", index)
	log.Fatal(http.ListenAndServe(":10000", r))
}

func main() {
	handleRequests()
}
