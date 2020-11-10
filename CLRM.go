package main

type CLRM struct {
	ols OLS
}

type ConfidenceInterval struct {
	start float32
	stop  float32
}

func (rm CLRM) calculateTStat(betaEst float32, betaEstSE float32, betaHypZero float32) float32 {
	return (betaEst - betaHypZero) / betaEstSE
}

func (rm CLRM) rejectTestOfSignificanceTestTwoSided(critValue float32, tStat float32) bool {
	lower := -critValue
	upper := critValue
	inNonRejectionZone := (lower < tStat) && (tStat < upper)
	return !inNonRejectionZone
}

func (rm CLRM) rejectTestOfSignificanceTestOneSided(critValue float32, tStat float32, leftTail bool) bool {
	if leftTail {
		return tStat <= critValue
	}
	return critValue <= tStat
}

func (rm CLRM) calculateConfidenceInterval(betaEst float32, betaEstSE float32, tCrit float32) ConfidenceInterval {
	locStart := betaEst - (tCrit * betaEstSE)
	locEnd := betaEst + (tCrit * betaEstSE)
	return ConfidenceInterval{
		start: locStart,
		stop:  locEnd,
	}
}

func (rm CLRM) rejectConfidenceIntervalTest(interval ConfidenceInterval, betaHypZero float32) bool {
	withinInterval := (interval.start <= betaHypZero) && (betaHypZero <= interval.stop)
	return !withinInterval
}
