package main

import (
	"fmt"
)

func main() {

	series := Dataseries{
		x: NumberSequence{5, 8, 9, 14, 17},
		y: NumberSequence{10, 12, 15, 20, 21},
	}

	meanX := series.x.calculateMean()
	varX := series.x.calculateVariance()
	sdX := series.x.calculateSd()

	fmt.Println(meanX)
	fmt.Println(varX)
	fmt.Println(sdX)

	cov, _ := series.calculateCoVariance()

	fmt.Println("Covariance ", cov)

}
