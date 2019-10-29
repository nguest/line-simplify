package algorithms

import (
	"fmt"
	"line-simplify/tracks"
	"math"
	"time"
)

func LeonardoOptimize(data []tracks.Datum) []tracks.Datum {
	defer timeTrack(time.Now(), "LeonardoOptimize")
	pnts := len(data)
	distance, maxDist, max2Dist := initDmVal(data)

	dMin, dMinI, dMinJ := initDMin(pnts, distance, maxDist)
	maxenddist, maxendpunkt, leaveout := initMaxEnd(pnts, distance, max2Dist)

	var max1, max2, max3, max4, max5, maxroute, bestfai, bestflach int
	var i1leaveout, fsleaveout, dreieckleaveout, flachleaveout, faileaveout int
	var max1flach, max2flach, max3flach, max4flach, max5flach, bflpdmc int
	var max1fai, max2fai, max3fai, max4fai, max5fai, baipdmc int

	max2d2 := max2Dist * 2
	max2d7 := max2Dist * 7
	max2d3 := max2Dist * 3
	i2cmp := pnts - 2

	var firstGuess func()
	firstGuess = func() {
		pnts := len(data) // TODO
		var distance []int
		var a, b, c, d, u, tmp int
		var bestfai, bestflach int
		/* geratene L�sung f�r freie Strecke: */
		/* max1 = 0;       /* Start ganz vorne */
		/* max2 = pnts/4;  /* Erste Wende nach einem 4tel der Strecke */
		/* max3 = pnts/2;  /* Zweite Wende etwa in der Mitte */
		/* max4 = pnts*3/4;/* Dritte Wende etwa nach 3/4teln der Strecke */
		/* max5 = pnts-1;  /* Endpunkt ganz hinten */

		max1 := pnts * 3 / 8 /* Start ganz vorne */
		max2 := pnts * 4 / 8 /* Erste Wende nach einem 4tel der Strecke */
		max3 := pnts * 5 / 8 /* Zweite Wende etwa in der Mitte */
		max4 := pnts * 6 / 8 /* Dritte Wende etwa nach 3/4teln der Strecke */
		max5 := pnts - 1     /* Endpunkt ganz hinten */
		maxroute = distance[max1+pnts*max2] + distance[max2+pnts*max3] + distance[max3+pnts*max4] + distance[max4+pnts*max5]

		/* geratene L�sung f�r ein Dreieck begonnen auf einem Schenkel */
		i1 := 0            /* Start ganz vorne */
		i2 := pnts / 6     /* Erste Wende nach einem Sechstel */
		i3 := pnts / 2     /* Zweite Wende in der Mitte */
		i4 := pnts * 5 / 6 /* Dritte Wende */
		i5 := pnts - 1     /* Endpunkt ganz hinten */
		a = distance[i2+pnts*i3]
		b = distance[i3+pnts*i4]
		c = distance[i2+pnts*i4]
		d = distance[i1+pnts*i5]

		u = a + b + c

		if d*5 <= u { /* zuf�llig ein Dreieck gefunden? */
			tmp = u * 7
			if (a*25 >= tmp) && (b*25 >= tmp) && (c*25 >= tmp) { /* zuf�llig FAI-D gefunden? */
				bestfai = u - d
				max1fai = i1
				max2fai = i2
				max3fai = i3
				max4fai = i4
				max5fai = i5
			} else { /* Flaches Dreieck */
				bestflach = u - d
				max1flach = i1
				max2flach = i2
				max3flach = i3
				max4flach = i4
				max5flach = i5
			}
		}
		firstGuess()

		/* geratene L�sung f�r eine Dreieck begonnen an erster Wende */
		i1 = 0
		i2 = 0            /* Start und erste Wende ganz vorne */
		i3 = pnts / 3     /* zweite Wende nach 1/3 der Strecke */
		i4 = pnts * 2 / 3 /* dritte Wende nach 2/3 der Strecke */
		i5 = pnts - 1     /* Endpunkt ganz hinten */
		a = distance[i2+pnts*i3]
		b = distance[i3+pnts*i4]
		c = distance[i2+pnts*i4]
		d = distance[i1+pnts*i5]
		u = a + b + c
		if d*5 <= u { /* zuf�llig ein Dreieck gefunden? */
			tmp = u * 7
			if (a*25 >= tmp) && (b*25 >= tmp) && (c*25 >= tmp) { /* zuf�llig FAI-D gefunden? */
				if (u - d) > bestfai {
					bestfai = u - d
					max1fai = i1
					max2fai = i2
					max3fai = i3
					max4fai = i4
					max5fai = i5
				}
			} else { /* Flaches Dreieck */
				if (u - d) > bestflach {
					bestflach = u - d
					max1flach = i1
					max2flach = i2
					max3flach = i3
					max4flach = i4
					max5flach = i5
				}
			}
		}
	}

	fmt.Println("calculating best waypoints.. for more than 500 points need some minutes, press Ctrl-C for intermediate results...")

	for i2 := 0; i2 < i2cmp; i2++ { /* 1.Wende */ /* i1leaveout = 1; kann wech */
		e := 0
		i1 := 0
		for i := 0; i < i2; i += i1leaveout { /* Starting point for free distance is separately optimized  */
			tmp := distance[i+pnts*i2]
			if tmp >= e {
				e = tmp
				i1 = i
			}
			i1leaveout = 1
			/*  MANOLIS if ((i1leaveout=(e-tmp)/max2dist)<1) i1leaveout = 1; */
		} /* e, i1 enthalten fuer dieses i2 den besten Wert  e, i1 contain the best value for this i2  */

		mrme := maxroute - e
		i4cmp := i2 + 2

		for i4 := pnts - 1; i4 >= i4cmp; i4 -= leaveout { /* 3.Wende von hinten optimieren */
			c := distance[i2+pnts*i4]
			c25 := c * 25
			d := fdMin(i2, i4, dMin)
			d5minusc := d*5 - c
			dmc := d - c
			bflpdmc = bestflach + dmc
			baipdmc = bestfai + dmc
			maxaplusb := 0 /* leaveout = 1;  eigentlich nicht notwendig */
			f := maxend(i4, maxenddist)
			mrmemf := mrme - f
			epf := e + f
			i3 := i2 + 1
			for i := i3; i < i4; i += leaveout { /* 2.Wende separat optimieren */
				a := distance[i2+pnts*i]
				b := distance[i+pnts*i4]
				aplusb := a + b
				if aplusb > maxaplusb { /* findet gr��tes a+b (und auch gr��tes Dreieck) */
					maxaplusb = aplusb
					i3 = i
				}
				if d5minusc <= aplusb { /* Dreieck gefunden 5*d<= a+b+c */
					u := aplusb + c
					tmp := u * 7
					if c25 >= tmp && a*25 >= tmp && b*25 >= tmp { /* FAI-D gefunden */
						w := u - d
						if w > bestfai { /* besseres FAI-D gefunden */
							max1fai = fdMinI(i2, i4, dMinI)
							max2fai = i2
							max3fai = i
							max4fai = i4
							max5fai = fdMinJ(i2, i4, dMinJ)
							bestfai = w
							baipdmc = w + dmc
						}
					} else { /* nicht FAI=flaches Dreieck gefunden */
						w := u - d
						if w > bestflach {
							max1flach = fdMinI(i2, i4, dMinI)
							max2flach = i2
							max3flach = i
							max4flach = i4
							max5flach = fdMinJ(i2, i4, dMinJ)
							bestflach = w
							bflpdmc = bestflach + dmc
						}
					}
				}
				/* leaveout = 1; */
				fsleaveout = (mrmemf-aplusb)/max2d2 + 1 /* +1 wg. > */
				/* if (fsleaveout>1) { */
				dreieckleaveout = (d5minusc - aplusb) / max2d2
				flachleaveout = (bflpdmc-aplusb)/max2d2 + 1 /* +1 wg > */
				faileaveout = (baipdmc-aplusb)/max2d2 + 1   /* +1 wg > */
				leaveout = MIN(flachleaveout, faileaveout)
				leaveout = MAX(leaveout, dreieckleaveout)
				leaveout = MIN(leaveout, fsleaveout)
				if leaveout < 1 {
					leaveout = 1
				}
				/* MANOLIS */
				leaveout = 1
				/*}*/
				/* printf("leaveouts: fs=%d dr=%d fl=%d fai=%d insgesamt=%d\n", fsleaveout,dreieckleaveout,flachleaveout,faileaveout,leaveout); */
			} /* maxaplusb, i3 enthalten fuer dieses i2 und i4 besten Wert */
			tmp := maxaplusb + epf
			if tmp > maxroute {
				max1 = i1
				max2 = i2
				max3 = i3
				max4 = i4
				max5 = maxendi(i4, maxendpunkt)
				maxroute = tmp
				mrme = tmp - e
			}
			/* leaveout = 1;*/
			fsleaveout = (mrmemf-maxaplusb)/max2d2 + 1 /* )>1) { */
			dreieckleaveout = (d5minusc - maxaplusb) / max2d7
			flachleaveout = (bflpdmc-maxaplusb)/max2d3 + 1
			faileaveout = (baipdmc-maxaplusb)/max2d3 + 1
			leaveout = MIN(flachleaveout, faileaveout)
			leaveout = MAX(leaveout, dreieckleaveout)
			leaveout = MIN(leaveout, fsleaveout)
			if leaveout < 1 {
				leaveout = 1
			}
		}
	}
	//printbest(max1, max2, max3, max4, max5, maxroute, bestfai, bestflach, pnts, distance)

	freeFlightKm := float64(maxroute) / 1000.0
	freeTriangleKm := float64(bestflach) / 1000.0
	FAITriangleKm := float64(bestfai) / 1000.0

	freeFlightPoints := freeFlightKm * 1.5
	freeTrianglePoints := freeTriangleKm * 1.75
	FAITrianglePoints := FAITriangleKm * 2.0

	if freeFlightPoints > freeTrianglePoints && freeFlightPoints > FAITrianglePoints {
		fmt.Println("OUT BEST_FLIGHT_TYPE FREE_FLIGHT")
	} else if freeTrianglePoints > FAITrianglePoints {
		/*
		 * Die Dreiecke bestehen aus den Schenkeln a, b und c. Von dieser Strecke
		 * wird die Distanz d zwischen Start- und Endpunkt abgezogen
		 */
		fmt.Println("OUT BEST_FLIGHT_TYPE FREE_TRIANGLE")
	} else {
		fmt.Println("OUT BEST_FLIGHT_TYPE FAI_TRIANGLE")
	}

	/* Print all opti results          */
	fmt.Println("OUT TYPE FREE_FLIGHT")
	fmt.Printf("OUT FLIGHT_KM %f\n", freeFlightKm)
	fmt.Printf("OUT FLIGHT_POINTS %f\n", freeFlightPoints)

	fmt.Printf("DEBUG Best free Flight: %f km = %f Points\n", freeFlightKm, freeFlightPoints)
	fmt.Println("OUT ")
	printpoint(max1)
	fmt.Printf("\n")
	fmt.Printf("OUT ")
	printpoint(max2)
	fmt.Printf(" %f km\n", float64(distance[max1+pnts*max2])/1000.0)
	fmt.Printf("OUT ")
	printpoint(max3)
	fmt.Printf(" %f km\n", float64(distance[max2+pnts*max3])/1000.0)
	fmt.Printf("OUT ")
	printpoint(max4)
	fmt.Printf(" %f km\n", float64(distance[max3+pnts*max4])/1000.0)
	fmt.Printf("OUT ")
	printpoint(max5)
	fmt.Printf(" %f km\n", float64(distance[max4+pnts*max5])/1000.0)

	fmt.Printf("OUT TYPE FREE_TRIANGLE\n")
	fmt.Printf("OUT FLIGHT_KM %f\n", freeTriangleKm)
	fmt.Printf("OUT FLIGHT_POINTS %f\n", freeTrianglePoints)

	fmt.Printf("DEBUG Best free Triangle: %f km = %f Points\n", float64(bestflach)/1000.0, float64(bestflach)/1000.0*1.75)
	fmt.Printf("OUT ")

	printpoint(max1flach)
	fmt.Printf("\n")
	fmt.Printf("OUT ")
	printpoint(max2flach)
	fmt.Printf(" %f km=d\n", float64(distance[max1flach+pnts*max5flach])/1000.0)
	fmt.Printf("OUT ")
	printpoint(max3flach)
	fmt.Printf(" %f km=a\n", float64(distance[max2flach+pnts*max3flach])/1000.0)
	fmt.Printf("OUT ")
	printpoint(max4flach)
	fmt.Printf(" %f km=b\n", float64(distance[max3flach+pnts*max4flach])/1000.0)
	fmt.Printf("OUT ")
	printpoint(max5flach)
	fmt.Printf(" %f km=c\n", float64(distance[max2flach+pnts*max4flach])/1000.0)

	fmt.Printf("OUT TYPE FAI_TRIANGLE\n")
	fmt.Printf("OUT FLIGHT_KM %f\n", FAITriangleKm)
	fmt.Printf("OUT FLIGHT_POINTS %f\n", FAITrianglePoints)

	fmt.Printf("bestes FAI Dreieck: %f km = %f Punkte\n", float64(bestfai)/1000.0, float64(bestfai)/1000.0*2.0)
	fmt.Printf("OUT ")
	printpoint(max1fai)
	fmt.Printf("\n")
	fmt.Printf("OUT ")
	printpoint(max2fai)
	fmt.Printf(" %f km=d\n", float64(distance[max1fai+pnts*max5fai])/1000.0)
	fmt.Printf("OUT ")
	printpoint(max3fai)
	fmt.Printf(" %f km=a\n", float64(distance[max2fai+pnts*max3fai])/1000.0)
	fmt.Printf("OUT ")
	printpoint(max4fai)
	fmt.Printf(" %f km=b\n", float64(distance[max3fai+pnts*max4fai])/1000.0)
	fmt.Printf("OUT ")
	printpoint(max5fai)
	fmt.Printf(" %f km=c\n", float64(distance[max2fai+pnts*max4fai])/1000.0)

	freeFlight := []tracks.Datum{
		data[max1],
		data[max2],
		data[max3],
		data[max4],
		data[max5],
	}
	return freeFlight
}

type PointList struct {
	tracks.Datum
	next tracks.Datum
}

// initialize distance matrix values
func initDmVal(data []tracks.Datum) ([]int, int, int) {
	fmt.Println("Initializing distance matrix values...")
	dFak := 6371000.0
	piDiv180 := math.Pi / 180.0
	pnts := len(data)

	lonPnts := make([]float64, pnts)
	latPnts := make([]float64, pnts)
	lonRad := make([]float64, pnts)
	latRad := make([]float64, pnts)
	sinLat := make([]float64, pnts)
	cosLat := make([]float64, pnts)
	distance := make([]int, pnts*pnts)

	// build the weird pointList
	pointList := make([]PointList, pnts)
	for i := 0; i < pnts-1; i++ {
		pointList[i].Datum = data[i]
		pointList[i].next = data[i+1]
	}
	cmp := pnts + 1

	for i := 0; i < pnts; i++ {
		lonPnts[i] = data[i].Lon
		lonRad[i] = data[i].Lon * piDiv180
		latPnts[i] = data[i].Lat
		latRad[i] = data[i].Lat * piDiv180

		sinLat[i] = math.Sin(latRad[i])
		cosLat[i] = math.Cos(latRad[i])
		//pointList = pointList.next

	}

	maxDist := 0  /* recalculate the maximum distance between any two points */
	max2Dist := 0 /* recalculate the maximum distance between two consecutive points */
	maxTDist := 0 /* max takeoff distance */

	cmp = pnts - 1 /* Memorize loop comparison value for fast calculation */

	var maxp1, maxp2, max2p1, max2p2 int
	var maxT1, maxT2 int

	for i := 0; i < cmp; i++ { /* diese Schleife NICHT RÜCKWÄRTS!!! */
		sli := sinLat[i]
		cli := cosLat[i]
		lri := lonRad[i]

		j := i + 1

		dist := int(dFak*math.Acos(sli*sinLat[j]+cli*cosLat[j]*math.Cos(lri-lonRad[j])) + 0.5)
		distance[i+pnts*j] = dist
		if dist > max2Dist {
			max2p1 = i
			max2p2 = j
			max2Dist = dist /* weiteste Distanz merken */
		}

		/* compute max distnace from point 0 (takeoff */
		if distance[pnts*j] > maxTDist {
			maxT2 = i
			maxTDist = distance[pnts*j]
		}

		for j = i + 2; j < pnts; j++ { /* Durchlauf j=i+1 rausgezogen */
			dist := int(dFak*math.Acos(sli*sinLat[j]+cli*cosLat[j]*math.Cos(lri-lonRad[j])) + 0.5)
			distance[i+pnts*j] = dist
			if dist > maxDist {
				maxp1 = i
				maxp2 = j
				maxDist = dist /* ggf. weiteste Distanz merken */
			}
		}
		/* DEBUG if (i+1==j) { comparedistances(i,j); } */

	}
	if max2Dist > maxDist {
		maxDist = max2Dist
		maxp1 = max2p1
		maxp2 = max2p2
	}

	fmt.Printf("DEBUG maximal distance between any 2 points: %d meters\n", maxDist)
	fmt.Printf("OUT MAX_LINEAR_DISTANCE %d\n", maxDist)
	fmt.Printf("TOTAL TRACKLOG POINTS %d\n", pnts)
	fmt.Printf("DEBUG P1: %d\n", maxp1)
	fmt.Printf("DEBUG P2: %d\n", maxp2)

	fmt.Printf("OUT TYPE FreeFlight0TP\n==============================\n")
	fmt.Printf("OUT FLIGHT_KM %f\n", float64(maxDist)/1000.0)
	fmt.Printf("OUT FLIGHT_POINTS %f\n", float64(maxDist)/1000.0)
	fmt.Printf("OUT ")
	printpoint(maxp1)
	fmt.Printf("\n")
	fmt.Printf("OUT ")
	printpoint(maxp2)
	fmt.Printf(" %f km\n", float64(distance[maxp1+pnts*maxp2])/1000.0)

	fmt.Printf("OUT TYPE MaxTakeoffDistance\n==============================\n")
	fmt.Printf("OUT FLIGHT_KM %f\n", float64(maxTDist)/1000.0)
	fmt.Printf("OUT FLIGHT_POINTS %f\n", float64(maxTDist)/1000.0)
	fmt.Printf("OUT ")
	printpoint(maxT1)
	fmt.Printf("\n")
	fmt.Printf("OUT ")
	printpoint(maxT2)
	fmt.Printf(" %f km\n", float64(distance[maxT1+pnts*maxT2])/1000.0)

	// fmt.Printf("DEBUG START_TIME %d\n", timepnts[0])
	// fmt.Printf("DEBUG END_TIME %d\n", timepnts[pnts-1])
	// duration = timepnts[pnts-1]- timepnts[0]
	// fmt.Printf("DEBUG DURATION_SEC %d\n", duration)
	// fmt.Printf("DEBUG DURATION %2d:%2d:%2d\n", duration/3600, (duration%3600)/60, duration%60)
	fmt.Printf("DEBUG maximal distance between 2 successive points: %d meters\n", max2Dist)
	return distance, maxDist, max2Dist
}

func initDMin(pnts int, distance []int, maxDist int) ([][]int, [][]int, [][]int) {

	dMin := make([][]int, pnts)
	for i := range dMin {
		dMin[i] = make([]int, pnts)
	}
	dMinI := make([][]int, pnts)
	for i := range dMinI {
		dMinI[i] = make([]int, pnts)
	}
	dMinJ := make([][]int, pnts)
	for i := range dMinJ {
		dMinJ[i] = make([]int, pnts)
	}
	//var dMinIndex []int

	var i, j, d, mini, minj int
	minimum := maxDist
	fmt.Println("initializing dmin(i,j) with best start/endpoints for triangles...")

	for j = pnts - 1; j > 0; j-- { /* erste Zeile separat behandeln: treat first line separately */
		d := distance[0+pnts*j]
		if d < minimum { /* d <=; minimum if equivalent point is to be found farther in the track */
			minimum = d
			minj = j
		}
		//dMin[] = minimum
		dMinI[0][j] = 0
		dMinJ[0][j] = minj

	}

	for i = 1; i < pnts-1; i++ { /* folgenden Zeilen von vorheriger ableiten:
		derive the following lines from previous ones */
		j = pnts - 1 /* letzte Spalte zur Initialisierung des Minimums getrennt behandeln : treat the last column separately to initialize the minimum*/
		minimum = dMin[i-1][j]
		mini = dMinI[i-1][j]
		minj = dMinJ[i-1][j]

		d = distance[i+pnts*j]

		if d < minimum {
			minimum = d
			mini = i
			minj = j
		}
		dMin[i][j] = minimum
		dMinI[i][j] = mini
		dMinJ[i][j] = minj

		for j := pnts - 2; j > i; j-- { /* andere spalten von hinten nach vorne bearbeiten */
			d = distance[i+pnts*j]
			if d < minimum { /* aktueller Punkt besser als bisheriges Minimum? */
				/* d<=minimum falls gleichwertiger Punkt weiter vorne im track gefunden werden soll */
				minimum = d
				mini = i
				minj = j
			}
			d = dMin[i-1][j]
			if d < minimum { /* Minimum aus vorheriger Zeile besser? : Minimum from previous line better? */
				minimum = d
				mini = dMinI[i-1][j]
				minj = dMinJ[i-1][j]
			}
			dMin[i][j] = minimum
			dMinI[i][j] = mini
			dMinJ[i][j] = minj
		}
	}
	return dMin, dMinI, dMinJ
}

func initMaxEnd(pnts int, distance []int, max2Dist int) ([]int, []int, int) {
	//	pnts := len(data)  // TODO dry
	//	max2Dist := 0      // TODO pass this VALUE!!!
	//	var distance []int // TODO pass this VALUE!!!

	maxenddist := make([]int, pnts)
	maxendpunkt := make([]int, pnts)

	var w3, i, f, maxf, besti, leaveout int
	fmt.Println("initializing maxenddist[] with maximal distance to best endpoint ...")

	for w3 = pnts - 1; w3 > 1; w3-- {
		maxf = 0
		leaveout = 1
		besti = pnts - 1
		for i = pnts - 1; i >= w3; i -= leaveout {
			f = distance[w3+pnts*i]
			if f >= maxf {
				maxf = f
				besti = i
			}
			leaveout = (maxf - f) / max2Dist
			if leaveout < 1 {
				leaveout = 1
			}
		}
		maxenddist[w3] = maxf
		maxendpunkt[w3] = besti
	}
	return maxenddist, maxendpunkt, leaveout
}

// func firstguess(data []tracks.Datum) int {
// 	pnts := len(data) // TODO
// 	var distance []int
// 	var a, b, c, d, u, tmp int
// 	var bestfai, bestflach int
// 	/* geratene L�sung f�r freie Strecke: */
// 	/* max1 = 0;       /* Start ganz vorne */
// 	/* max2 = pnts/4;  /* Erste Wende nach einem 4tel der Strecke */
// 	/* max3 = pnts/2;  /* Zweite Wende etwa in der Mitte */
// 	/* max4 = pnts*3/4;/* Dritte Wende etwa nach 3/4teln der Strecke */
// 	/* max5 = pnts-1;  /* Endpunkt ganz hinten */

// 	max1 := pnts * 3 / 8 /* Start ganz vorne */
// 	max2 := pnts * 4 / 8 /* Erste Wende nach einem 4tel der Strecke */
// 	max3 := pnts * 5 / 8 /* Zweite Wende etwa in der Mitte */
// 	max4 := pnts * 6 / 8 /* Dritte Wende etwa nach 3/4teln der Strecke */
// 	max5 := pnts - 1     /* Endpunkt ganz hinten */
// 	maxroute := distance[max1+pnts*max2] + distance[max2+pnts*max3] + distance[max3+pnts*max4] + distance[max4+pnts*max5]

// 	/* geratene L�sung f�r ein Dreieck begonnen auf einem Schenkel */
// 	i1 := 0            /* Start ganz vorne */
// 	i2 := pnts / 6     /* Erste Wende nach einem Sechstel */
// 	i3 := pnts / 2     /* Zweite Wende in der Mitte */
// 	i4 := pnts * 5 / 6 /* Dritte Wende */
// 	i5 := pnts - 1     /* Endpunkt ganz hinten */
// 	a = distance[i2+pnts*i3]
// 	b = distance[i3+pnts*i4]
// 	c = distance[i2+pnts*i4]
// 	d = distance[i1+pnts*i5]

// 	u = a + b + c

// 	if d*5 <= u { /* zuf�llig ein Dreieck gefunden? */
// 		tmp = u * 7
// 		if (a*25 >= tmp) && (b*25 >= tmp) && (c*25 >= tmp) { /* zuf�llig FAI-D gefunden? */
// 			bestfai := u - d
// 			max1fai := i1
// 			max2fai := i2
// 			max3fai := i3
// 			max4fai := i4
// 			max5fai := i5
// 		} else { /* Flaches Dreieck */
// 			bestflach := u - d
// 			max1flach := i1
// 			max2flach := i2
// 			max3flach := i3
// 			max4flach := i4
// 			max5flach := i5
// 		}
// 		//return FAI, flach
// 	}

// 	/* geratene L�sung f�r eine Dreieck begonnen an erster Wende */
// 	i1 = 0
// 	i2 = 0            /* Start und erste Wende ganz vorne */
// 	i3 = pnts / 3     /* zweite Wende nach 1/3 der Strecke */
// 	i4 = pnts * 2 / 3 /* dritte Wende nach 2/3 der Strecke */
// 	i5 = pnts - 1     /* Endpunkt ganz hinten */
// 	a = distance[i2+pnts*i3]
// 	b = distance[i3+pnts*i4]
// 	c = distance[i2+pnts*i4]
// 	d = distance[i1+pnts*i5]
// 	u = a + b + c
// 	if d*5 <= u { /* zuf�llig ein Dreieck gefunden? */
// 		tmp = u * 7
// 		if (a*25 >= tmp) && (b*25 >= tmp) && (c*25 >= tmp) { /* zuf�llig FAI-D gefunden? */
// 			if (u - d) > bestfai {
// 				bestfai := u - d
// 				max1fai := i1
// 				max2fai := i2
// 				max3fai := i3
// 				max4fai := i4
// 				max5fai := i5
// 			}
// 		} else { /* Flaches Dreieck */
// 			if (u - d) > bestflach {
// 				bestflach := u - d
// 				max1flach := i1
// 				max2flach := i2
// 				max3flach := i3
// 				max4flach := i4
// 				max5flach := i5
// 			}
// 		}
// 	}
// 	return maxroute
// }

func MIN(x, y int) int {
	if x < y {
		return x
	}
	return y
}
func MAX(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// setter fns (I think)
func dMin(v, x, y, pnts int, distance []int) {
	distance[y+pnts*x] = v
}
func dMinI(v, x, y, pnts int, dMinIndex []int) {
	dMinIndex[x+pnts*y] = v
}
func dMinJ(v, x, y, pnts int, dMinIndex []int) {
	dMinIndex[y+pnts*x] = v
}

// getter fns
func fdMin(x, y int, dMin [][]int) int {
	return dMin[x][y]
}
func fdMinI(x, y int, dMinI [][]int) int {
	return dMinI[x][y]
}
func fdMinJ(x, y int, dMinJ [][]int) int {
	return dMinJ[x][y]
}

func maxend(v int, maxenddist []int) int {
	return maxenddist[(v)]
}
func maxendi(v int, maxendpunkt []int) int {
	return maxendpunkt[(v)]
}

/*
* Best Score for free distance, flat triangle and FAI triangle
* spend exactly with total distance in km and score on 1000stel
* with the OLC points are rounded on 100stel, more?ber rounding
* No statement is made stages, which seems OLC server to however already round
* stages on 100stel km = dezimeter.
* Nevertheless in meters and not in dezimetern one counts here,
* comes there it otherwise to rounding errors.
 */

// func printbest(max1, max2, max3, max4, max5, maxroute, bestfai, bestflach, pnts int, distance []int) {

// }

func printpoint(i int) {
	fmt.Printf("INDEX %v", i)
}
