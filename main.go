package main

import "fmt"

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

	if counter != 58 {
		panic(fmt.Errorf("GetByFips(%s) yeilded %d structures; expected 58", fips, counter))
	} else {
		fmt.Printf("GetByFips(%s) yeilded %d structures; expected 58\n", fips, counter)
	}
	if population != 115 {
		panic(fmt.Errorf("GetByFips(%s) yeilded population of %d across all structures; expected 115", fips, population))
	} else {
		fmt.Printf("GetByFips(%s) yeilded population of %d across all structures; expected 115\n", fips, population)
	}
	if totVal != 44632201.8453 {
		panic(fmt.Errorf("GetByFips(%s) yeilded total value of %f across all structures; expected 44632201", fips, totVal))
	} else {
		fmt.Printf("GetByFips(%s) yeilded total value of %f across all structures; expected 44632201\n", fips, totVal)
	}
}
