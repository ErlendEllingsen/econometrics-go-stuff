package main

type FTest struct {
	reg OLSMultivariate
}

func (ft FTest) calculateTestStat(rrss float32, urss float32, m float32, T, k float32) float32 {
	return ((rrss - urss) / m) / (urss / (T - k))
}

// TODO: Implement me
