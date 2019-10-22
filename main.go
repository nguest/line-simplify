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

	//out := DPByCount(in.Data, 5)
	out := Visvalingam(in.Data, 5)

	var X struct {
		In      [][]float64 `json:"in"`
		Out     [][]float64 `json:"out"`
		OutDist float64     `json:"outDist"`
	}
	for _, v := range out {
		X.Out = append(X.Out, []float64{v.Lon, v.Lat})
	}
	for _, v := range in.Data {
		X.In = append(X.In, []float64{v.Lon, v.Lat})
	}
	X.OutDist = GetTotalTrackLength(out)
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
