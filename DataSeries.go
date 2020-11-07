package main

import (
	"errors"
)

type Dataseries struct {
	y NumberSequence
	x NumberSequence
}

func (d Dataseries) calculateCoVariance() (float32, error) {
	if len(d.x) != len(d.y) {
		return -1, errors.New("Unable to calculate covariance if slice length of both variables not equal")
	}

	meanX := d.x.calculateMean()
	meanY := d.y.calculateMean()

	// Start building the mathematics formula
	numerator := float32(0) // to be calclated
	denominator := float32(len(d.x))

	for i := 0; i < int(denominator); i++ {
		elemX := float32(d.x[i])
		elemY := float32(d.y[i])
		numerator += ((elemX - meanX) * (elemY - meanY))
	}

	covariance := numerator / denominator // implementing COVARIANCE.P
	return covariance, nil
}

func (d Dataseries) calculateCorrelation() (float32, error) {
	cov, err := d.calculateCoVariance()
	if err != nil {
		return -1, errors.New("Unable to calculate correlation if slice length of both variables not equal")
	}
	corr := (cov / (d.x.calculateSd() * d.y.calculateSd()))
	return corr, nil // implementing CORREL
}
