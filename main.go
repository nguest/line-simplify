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
	out := DouglasPeucker(in.Data, 0.02)

	var X struct {
		In  [][]float64 `json:"in"`
		Out [][]float64 `json:"out"`
	}
	for _, v := range out {
		X.Out = append(X.Out, []float64{v.Lon, v.Lat})
	}
	for _, v := range in.Data {
		X.In = append(X.In, []float64{v.Lon, v.Lat})
	}
	json.Marshal(X)
	t.Execute(w, X)
	//X := "hello"
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

	// fmt.Println("finalRESULT", DouglasPeucker([]Datum{{0, 0}, {1, 0.1}, {2, -0.1},
	// 	{3, 5}, {4, 6}, {5, 7}, {6, 8.1}, {5, 9}, {8, 9}, {9, 9}, {10, 8}, {11, 8.5}}, 1))
}
