package main

import "math"

type NumberSequence []int32

func (ns NumberSequence) calculateMean() float32 {
	sum := int32(0)
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
	return variance / float32(len(ns)) // Implementing VAR.P (and thus STDEV.P)
}

func (ns NumberSequence) calculateSd() float32 {
	variance := ns.calculateVariance()
	return float32(math.Sqrt(float64(variance)))
}
