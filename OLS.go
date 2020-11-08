package main

import (
	"errors"
)

type OLS struct {
	ds Dataseries
}

func (ols OLS) estimateSlope() (float32, error) {
	d := ols.ds

	if len(d.x) != len(d.y) {
		return -1, errors.New("Unable to calculate ols slope valuee if slice length of both variables not equal")
	}

	meanX := d.x.calculateMean()
	meanY := d.y.calculateMean()

	// Start building the mathematics formula
	numerator := float32(0)   // to be calclated
	denominator := float32(0) // to be calclated

	for i := 0; i < len(d.x); i++ {
		elemX := float32(d.x[i])
		elemY := float32(d.y[i])
		numerator += ((elemX - meanX) * (elemY - meanY))
		denominator += ((elemX - meanX) * (elemX - meanX))
	}
	slopeCoefficient := numerator / denominator
	return slopeCoefficient, nil
}

func (ols OLS) estimateIntercept() (float32, error) {
	d := ols.ds

	if len(d.x) != len(d.y) {
		return -1, errors.New("Unable to calculate ols intercept if slice length of both variables not equal")
	}

	estimatedSlopeCoefficient, _ := ols.estimateSlope()

	meanX := d.x.calculateMean()
	meanY := d.y.calculateMean()

	interceptCofficient := meanY - (estimatedSlopeCoefficient * meanX)
	return interceptCofficient, nil
}
