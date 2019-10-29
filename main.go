package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"line-simplify/algorithms"
	"line-simplify/tracks"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func index(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println(err)
	}
	in, err := tracks.ReadAndParse("sampleTracks/sample1.igc")

	pCount := 5

	outDP := algorithms.DPByCount(in.Data, pCount)
	outVis := algorithms.Visvalingam(in.Data, pCount)
	outBrute := algorithms.LeonardoOptimize(in.Data)

	type L struct {
		Line  [][]float64 `json:"line"`
		Dist  float64     `json:"dist"`
		Title string      `json:"title"`
	}

	var X struct {
		In       L `json:"in"`
		OutDP    L `json:"outDP"`
		OutVis   L `json:"outVis"`
		OutBrute L `json:"outBrute"`
	}

	for _, v := range outDP {
		X.OutDP.Line = append(X.OutDP.Line, []float64{v.Lon, v.Lat})
	}
	X.OutDP.Title = "Douglas-Peucker"
	X.OutDP.Dist = tracks.GetTotalTrackLength(outDP)

	for _, v := range outVis {
		X.OutVis.Line = append(X.OutVis.Line, []float64{v.Lon, v.Lat})
	}
	X.OutVis.Title = "Visvalingam"
	X.OutVis.Dist = tracks.GetTotalTrackLength(outVis)

	for _, v := range outBrute {
		X.OutBrute.Line = append(X.OutBrute.Line, []float64{v.Lon, v.Lat})
	}
	X.OutBrute.Title = "Brute Force"
	X.OutBrute.Dist = tracks.GetTotalTrackLength(outBrute)

	for _, v := range in.Data {
		X.In.Line = append(X.In.Line, []float64{v.Lon, v.Lat})
	}
	X.In.Title = "Data In"
	X.In.Dist = tracks.GetTotalTrackLength(in.Data)

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
