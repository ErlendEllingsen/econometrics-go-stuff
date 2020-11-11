package main

type FTest struct {
	reg OLSMultivariate
}

func (ft FTest) calculateTestStat(rrss float64, urss float64, m float64, T, k float64) float64 {
	return ((rrss - urss) / m) / (urss / (T - k))
}


func (ft FTest) rejectFTest(critValue float64, fTestStat float64) bool {
	return fTestStat > critValue
}
