package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

// TrackData is a basic structure of all the data saved for a track.
type TrackData struct {
	Data             []Datum   `json:"data"`
	ID               string    `json:"id"`
	ElapsedTime      int       `json:"elapsedTime"`
	IsComplete       bool      `json:"isComplete"`
	StartEndDistance float64   `json:"startEndDistance"`
	Date             time.Time `json:"date"`
	//BoundingBox      BoundingBox `json:"boundingBox"`
}

func readLocalFile(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened", filepath)
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// ReadAndParse opens a file from local source, and parses its contents.
func ReadAndParse(filepath string) (TrackData, error) {
	var lines []string
	lines, _ = readLocalFile(filepath)
	trackData, err := parseIGC(lines)
	return trackData, err
}

// parseIgc takes an array of linestrings and converts to trackData object.
func parseIGC(lines []string) (TrackData, error) {
	var data []Datum
	var trackData TrackData
	var rawDate string

	for _, line := range lines {

		// parse the current date (DDMMYY)
		if len(line) > 5 && line[0:5] == "HFDTE" {
			if line[5:10] == "DATE:" {
				rawDate = line[10:16]
			} else {
				rawDate = line[5:11]
			}
		}

		// parse the B (fix) records
		if len(line) > 2 && line[:1] == "B" {

			var rawTimeOfDay, latHemi, lonHemi, validity string
			var latDeg, latMin, latFrac, lonDeg, lonMin, lonFrac, baroAlt, gpsAlt float64

			fmt.Sscanf(
				line,
				"B%6s%2f%2f%3f%1s%3f%2f%3f%1s%1s%5f%5f",
				&rawTimeOfDay, &latDeg, &latMin, &latFrac, &latHemi,
				&lonDeg, &lonMin, &lonFrac, &lonHemi,
				&validity, &baroAlt, &gpsAlt)

			lonMultiplier := 1.0
			if lonHemi == "W" {
				lonMultiplier = -1.0
			}

			latMultiplier := 1.0
			if latHemi == "S" {
				latMultiplier = -1.0
			}
			longitude := lonMultiplier * (lonDeg + (lonMin*1000+lonFrac)/1000/60)
			latitude := latMultiplier * (latDeg + (latMin*1000+latFrac)/1000/60)

			dataItem := Datum{
				Lat: latitude,
				Lon: longitude,
				Alt: gpsAlt,
				Ts:  getTimeStamp(rawDate, rawTimeOfDay),
			}
			// fmt.Println("{", dataItem.Lon, ",", dataItem.Lat, "},")
			data = append(data, dataItem)
		}
	}

	trackData.StartEndDistance = GetTotalTrackLength(data)
	trackData.Data = data
	trackData.Date = data[0].Ts

	return trackData, nil
}

func getTimeStamp(rawDate string, rawTimeOfDay string) time.Time {
	rawTimeStamp := fmt.Sprintf("%s%s", rawDate, rawTimeOfDay)
	parsedTime, _ := time.Parse("020106150405", rawTimeStamp)
	//timeStamp := parsedTime.Format("Mon, 02 Jan 2006 15:04:05 -0700")
	return parsedTime
}

// HFDTE 311011
// HFDTEDATE: 250819,01

// B 210939 4531841N 12240920W A0000000018
// B 112840 4624845N 00806464E A0214502144005003001258179

// "B%2u%2u%2u%2u%2u%3u%1[NS]%3u%2u%3u%1[WE]%c%5d%5d",
// 		       &hours, &mins, &secs, &lat_deg, &lat_min, &lat_frac,
// 		       lat_hemi, &lon_deg, &lon_min, &lon_frac, lon_hemi,
// 		       &validity, &pres_alt, &gnss_alt) != 14) {
//     fatal(MYNAME ": fix (B) record parse error\n%s\n", ibuf);
