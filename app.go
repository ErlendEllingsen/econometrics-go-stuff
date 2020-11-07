package main

import (
	"fmt"
	"math"
)

type NumberSequence []uint32

func (ns NumberSequence) calculateMean() float32 {
	sum := uint32(0)
	for _, v := range ns {
		sum += v
	}
	mean := float32(sum) / float32(len(ns))
	return mean
}

func (ns NumberSequence) calculateVariance() float32 {

	mean := ns.calculateMean()

	variance := float32(0)

	for i := 0; i < len(ns); i++ {
		elem := ns[i]
		variance += float32(math.Pow(float64(float32(elem)-float32(mean)), 2))
	}
	return variance / float32(len(ns))
}

func (ns NumberSequence) calculateSd() float32 {
	variance := ns.calculateVariance()
	return float32(math.Sqrt(float64(variance)))
}

type Dataseries struct {
	y NumberSequence
	x NumberSequence
}

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

}
