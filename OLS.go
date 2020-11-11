package main

import (
	"errors"
	"math"
)

type OLS struct {
	ds Dataseries
}

func (ols OLS) estimateSlope() (float64, error) {
	d := ols.ds

	if len(d.x) != len(d.y) {
		return -1, errors.New("Unable to calculate ols slope valuee if slice length of both variables not equal")
	}

	meanX := d.x.calculateMean()
	meanY := d.y.calculateMean()

	// Start building the mathematics formula
	numerator := float64(0)   // to be calclated
	denominator := float64(0) // to be calclated

	for i := 0; i < len(d.x); i++ {
		elemX := float64(d.x[i])
		elemY := float64(d.y[i])
		numerator += ((elemX - meanX) * (elemY - meanY))
		denominator += ((elemX - meanX) * (elemX - meanX))
	}
	slopeCoefficient := numerator / denominator
	return slopeCoefficient, nil
}

func (ols OLS) estimateIntercept() (float64, error) {
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

func (ols OLS) estimateY(intercept float64, slope float64, x float64) float64 {
	return intercept + (slope * x)
}

func (ols OLS) calculateRSS(intercept float64, slope float64) (float64, error) {

	d := ols.ds

	if len(d.x) != len(d.y) {
		return -1, errors.New("Unable to calculate RSS if slice length of both variables not equal")
	}

	rss := float64(0)

	for i := 0; i < len(d.x); i++ {
		elemX := float64(d.x[i])
		elemY := float64(d.y[i])

		estimatedY := ols.estimateY(intercept, slope, elemX)
		squaredDist := float64(math.Pow(float64(elemY-estimatedY), 2))
		rss += squaredDist
	}

	return rss, nil

}

func (ols OLS) calculateStandardError(rss float64) float64 { // Calculates residual standard error (regression error)
	d := ols.ds
	T := float64(len(d.x)) // T = N (Num of obs)
	return float64(math.Sqrt(float64(rss / (T - float64(2)))))
}

func (ols OLS) calculateEstimatorStandardErrors(residualStandardError float64) (alphaSE float64, betaSE float64) {

	d := ols.ds
	T := len(d.x)

	// Calculate alpha SE
	meanX := d.x.calculateMean()
	xSumPowNominator := float64(0)
	xSumDenominator := float64(0)
	for i := 0; i < T; i++ {
		elemX := float64(d.x[i])
		xSumPowNominator += float64(math.Pow(float64(elemX), 2))
		xSumDenominator += float64(math.Pow(float64(elemX-meanX), 2))
	}

	// alpha SE. Se Brooks, Chris, Introductory Econometrics for Finance 4th edition p. 110 for formulas
	locAlphaSE := residualStandardError * float64(math.Sqrt(float64(xSumPowNominator/(float64(T)*xSumDenominator))))

	// Calculate beta SE, partially same formula, so re-using some calculated vars
	locBetaSE := residualStandardError * float64(math.Sqrt(float64(float64(1)/(xSumDenominator))))

	return locAlphaSE, locBetaSE
}

func (ols OLS) calculateRSquared(intercept float64, slope float64) (float64, error) {

	d := ols.ds

	if len(d.x) != len(d.y) {
		return -1, errors.New("Unable to calculate RSS if slice length of both variables not equal")
	}

	explainedVariation := float64(0)
	unexplainedVariation := float64(0)

	meanY := d.y.calculateMean()

	for i := 0; i < len(d.x); i++ {
		elemX := float64(d.x[i])
		elemY := float64(d.y[i])
		estimatedElemY := ols.estimateY(intercept, slope, elemX)
		explainedVariation += float64(math.Pow(float64(estimatedElemY-meanY), 2))
		unexplainedVariation += float64(math.Pow(float64(elemY-estimatedElemY), 2))
	}

	totalVariation := explainedVariation + unexplainedVariation

	rSquared := 1 - (unexplainedVariation / totalVariation)
	return rSquared, nil

}

func (ols OLS) calculateRSquaredAdjusted(rSquared float64) (float64, error) {

	d := ols.ds

	if len(d.x) != len(d.y) {
		return -1, errors.New("Unable to calculate RSS if slice length of both variables not equal")
	}

	T := float64(len(d.x)) // T = num obs
	k := float64(2)        // k = num est. coeff (2 for singular)

	rssAdj := 1 - (((T - 1) / (T - k)) * (1 - rSquared))
	return rssAdj, nil
}

func (ols OLS) calculateDF() float64 {
	T := float64(len(ols.ds.x))
	k := float64(2) // k = num est. coeff (2 for singular)
	return float64(T - k)
}
