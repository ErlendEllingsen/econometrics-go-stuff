package main

import (
	"math"
)

type TTest struct {
	ols OLS
}

type ConfidenceInterval struct {
	start float64
	stop  float64
}

func (tt TTest) calculateTStat(betaEst float64, betaEstSE float64, betaHypZero float64) float64 {
	return (betaEst - betaHypZero) / betaEstSE
}

func (tt TTest) rejectTestOfSignificanceTestTwoSided(critValue float64, tStat float64) bool {
	upper := float64(math.Abs(float64(critValue)))
	lower := upper * -1
	inNonRejectionZone := (lower < tStat) && (tStat < upper)
	return !inNonRejectionZone
}

func (tt TTest) rejectTestOfSignificanceTestOneSided(critValue float64, tStat float64, leftTail bool) bool {
	if leftTail {
		return tStat <= critValue
	}
	return critValue <= tStat
}

func (tt TTest) calculateConfidenceInterval(betaEst float64, betaEstSE float64, tCrit float64) ConfidenceInterval {
	tCritAbs := float64(math.Abs(float64(tCrit)))
	locStart := betaEst - (tCritAbs * betaEstSE)
	locEnd := betaEst + (tCritAbs * betaEstSE)
	return ConfidenceInterval{
		start: locStart,
		stop:  locEnd,
	}
}

func (tt TTest) rejectConfidenceIntervalTest(interval ConfidenceInterval, betaHypZero float64) bool {
	withinInterval := (interval.start <= betaHypZero) && (betaHypZero <= interval.stop)
	return !withinInterval
}
