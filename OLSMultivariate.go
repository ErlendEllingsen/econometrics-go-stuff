package main

type OLSMultivariate struct {
	ds DataseriesMulti
}

func (ols OLSMultivariate) estimateSlopeValues() (float64, error) {
	return -1, nil
}

func (ols OLSMultivariate) estimateIntercept() (float64, error) {
	return -1, nil
}

func (ols OLSMultivariate) estimateY(intercept float64, slopeValuesMultipliedByX []float64) float64 {
	return -1
}
