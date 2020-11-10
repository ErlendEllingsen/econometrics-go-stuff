package main

type OLSMultivariate struct {
	ds DataseriesMulti
}

func (ols OLSMultivariate) estimateSlopeValues() (float32, error) {
	return -1, nil
}

func (ols OLSMultivariate) estimateIntercept() (float32, error) {
	return -1, nil
}

func (ols OLSMultivariate) estimateY(intercept float32, slopeValuesMultipliedByX []float32) float32 {
	return -1
}
