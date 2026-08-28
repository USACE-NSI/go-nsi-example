package main

import (
	"fmt"
	"math"
)

func main() {
	nsiapi := InitNSIApi()
	fips := "15005" //Kalawao county (smallest county in the us by population)
	counter := 0
	var population int32
	var totVal float64
	nsiapi.ByFips(fips, func(s NsiFeature) {
		counter++
		population += s.Properties.Pop2amu65 + s.Properties.Pop2amo65
		totVal += s.Properties.StructVal + s.Properties.ContVal

	})

	if counter != 55 {
		panic(fmt.Errorf("GetByFips(%s) yeilded %d structures; expected 55", fips, counter))
	} else {
		fmt.Printf("GetByFips(%s) yeilded %d structures; expected 55\n", fips, counter)
	}
	if population != 76 {
		panic(fmt.Errorf("GetByFips(%s) yeilded population of %d across all structures; expected 76", fips, population))
	} else {
		fmt.Printf("GetByFips(%s) yeilded population of %d across all structures; expected 76\n", fips, population)
	}
	diff := math.Abs(totVal - 25607115.933025)
	if diff > 1.0 {
		panic(fmt.Errorf("GetByFips(%s) yeilded total value of %f across all structures; expected 25607115.93", fips, totVal))
	} else {
		fmt.Printf("GetByFips(%s) yeilded total value of %f across all structures; expected 25607115.93\n", fips, totVal)
	}
}
