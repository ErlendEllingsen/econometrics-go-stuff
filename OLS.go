package main

import (
	"errors"
	"math"
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

func (ols OLS) estimateY(intercept float32, slope float32, x float32) float32 {
	return intercept + (slope * x)
}

func (ols OLS) calculateRSS(intercept float32, slope float32) (float32, error) {

	d := ols.ds

	if len(d.x) != len(d.y) {
		return -1, errors.New("Unable to calculate RSS if slice length of both variables not equal")
	}

	rss := float32(0)

	for i := 0; i < len(d.x); i++ {
		elemX := float32(d.x[i])
		elemY := float32(d.y[i])

		estimatedY := ols.estimateY(intercept, slope, elemX)
		squaredDist := float32(math.Pow(float64(elemY-estimatedY), 2))
		rss += squaredDist
	}

	return rss, nil

}

func (ols OLS) calculateStandardError(rss float32) float32 { // Calculates residual standard error (regression error)
	d := ols.ds
	T := float32(len(d.x)) // T = N (Num of obs)
	return float32(math.Sqrt(float64(rss / (T - float32(2)))))
}

func (ols OLS) calculateRSquared(intercept float32, slope float32) (float32, error) {

	d := ols.ds

	if len(d.x) != len(d.y) {
		return -1, errors.New("Unable to calculate RSS if slice length of both variables not equal")
	}

	explainedVariation := float32(0)
	unexplainedVariation := float32(0)

	meanY := d.y.calculateMean()

	for i := 0; i < len(d.x); i++ {
		elemX := float32(d.x[i])
		elemY := float32(d.y[i])
		estimatedElemY := ols.estimateY(intercept, slope, elemX)
		explainedVariation += float32(math.Pow(float64(estimatedElemY-meanY), 2))
		unexplainedVariation += float32(math.Pow(float64(elemY-estimatedElemY), 2))
	}

	totalVariation := explainedVariation + unexplainedVariation

	rSquared := 1 - (unexplainedVariation / totalVariation)
	return rSquared, nil

}

func (ols OLS) calculateRSquaredAdjusted(rSquared float32) (float32, error) {

	d := ols.ds

	if len(d.x) != len(d.y) {
		return -1, errors.New("Unable to calculate RSS if slice length of both variables not equal")
	}

	T := float32(len(d.x)) // T = num obs
	k := float32(2)        // k = num est. coeff (2 for singular)

	rssAdj := 1 - (((T - 1) / (T - k)) * (1 - rSquared))
	return rssAdj, nil
}

func (ols OLS) calculateDF() float32 {
	T := float32(len(ols.ds.x))
	k := float32(2) // k = num est. coeff (2 for singular)
	return float32(T - k)
}

func (ols OLS) calculateEstimatorStandardErrors(residualStandardError float32) (alphaSE float32, betaSE float32) {

	d := ols.ds
	T := len(d.x)

	// Calculate alpha SE
	meanX := d.x.calculateMean()
	xSumPowNominator := float32(0)
	xSumDenominator := float32(0)
	for i := 0; i < T; i++ {
		elemX := float32(d.x[i])
		xSumPowNominator += float32(math.Pow(float64(elemX), 2))
		xSumDenominator += float32(math.Pow(float64(elemX-meanX), 2))
	}

	// alpha SE. Se Brooks, Chris, Introductory Econometrics for Finance 4th edition p. 110 for formulas
	locAlphaSE := residualStandardError * float32(math.Sqrt(float64(xSumPowNominator/(float32(T)*xSumDenominator))))

	// Calculate beta SE, partially same formula, so re-using some calculated vars
	locBetaSE := residualStandardError * float32(math.Sqrt(float64(float32(1)/(xSumDenominator))))

	return locAlphaSE, locBetaSE
}
