package main

type CLRM struct {
	ols OLS
}

func (rm CLRM) calculateTStat(betaEst float32, betaEstSE float32, betaHypZero float32) float32 {
	return (betaEst - betaHypZero) / betaEstSE
}

func (rm CLRM) calculateConfidenceInterval(betaEst float32, betaEstSE float32, tCrit float32) (start float32, stop float32) {
	locStart := betaEst - (tCrit * betaEstSE)
	locEnd := betaEst + (tCrit * betaEstSE)
	return locStart, locEnd
}
