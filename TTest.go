package main

import (
	"math"
)

type TTest struct {
	ols OLS
}

type ConfidenceInterval struct {
	start float32
	stop  float32
}

func (tt TTest) calculateTStat(betaEst float32, betaEstSE float32, betaHypZero float32) float32 {
	return (betaEst - betaHypZero) / betaEstSE
}

func (tt TTest) rejectTestOfSignificanceTestTwoSided(critValue float32, tStat float32) bool {
	upper := float32(math.Abs(float64(critValue)))
	lower := upper * -1
	inNonRejectionZone := (lower < tStat) && (tStat < upper)
	return !inNonRejectionZone
}

func (tt TTest) rejectTestOfSignificanceTestOneSided(critValue float32, tStat float32, leftTail bool) bool {
	if leftTail {
		return tStat <= critValue
	}
	return critValue <= tStat
}

func (tt TTest) calculateConfidenceInterval(betaEst float32, betaEstSE float32, tCrit float32) ConfidenceInterval {
	tCritAbs := float32(math.Abs(float64(tCrit)))
	locStart := betaEst - (tCritAbs * betaEstSE)
	locEnd := betaEst + (tCritAbs * betaEstSE)
	return ConfidenceInterval{
		start: locStart,
		stop:  locEnd,
	}
}

func (tt TTest) rejectConfidenceIntervalTest(interval ConfidenceInterval, betaHypZero float32) bool {
	withinInterval := (interval.start <= betaHypZero) && (betaHypZero <= interval.stop)
	return !withinInterval
}
