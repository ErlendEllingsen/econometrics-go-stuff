package main

import "math"

type NumberSequence []float64

func (ns NumberSequence) calculateMean() float64 {
	sum := float64(0)
	for _, v := range ns {
		sum += v
	}
	mean := float64(sum) / float64(len(ns))
	return mean
}

func (ns NumberSequence) calculateVariance() float64 {

	mean := ns.calculateMean()

	variance := float64(0)

	for i := 0; i < len(ns); i++ {
		elem := ns[i]
		variance += float64(math.Pow(float64(float64(elem)-float64(mean)), 2))
	}
	return variance / float64(len(ns)) // Implementing VAR.P (and thus STDEV.P)
}

func (ns NumberSequence) calculateSd() float64 {
	variance := ns.calculateVariance()
	return float64(math.Sqrt(float64(variance)))
}
