package main

import (
	"fmt"
)

func main() {

	series := Dataseries{
		x: NumberSequence{5, 8, 9, 14, 17},
		y: NumberSequence{10, 12, 15, 20, 21},
	}

	meanX := series.x.calculateMean()
	varX := series.x.calculateVariance()
	sdX := series.x.calculateSd()

	fmt.Println(meanX)
	fmt.Println(varX)
	fmt.Println(sdX)

	cov, _ := series.calculateCoVariance()
	corr, _ := series.calculateCorrelation()

	fmt.Println("Covariance ", cov)
	fmt.Println("Correlation ", corr)

	// OLS
	ols := OLS{
		ds: series,
	}

	slope, _ := ols.estimateSlope()
	intercept, _ := ols.estimateIntercept()
	rss, _ := ols.calculateRSS(intercept, slope)
	residualSe := ols.calculateStandardError(rss)
	alphaSe, betaSe := ols.calculateEstimatorStandardErrors(residualSe)
	rSquared, _ := ols.calculateRSquared(intercept, slope)
	rSquaredAdj, _ := ols.calculateRSquaredAdjusted(rSquared)
	dF := ols.calculateDF()

	fmt.Println("OLS slope", slope)
	fmt.Println("OLS intercept", intercept)

	fmt.Println("OLS RSS", rss)
	fmt.Println("OLS SE Residual", residualSe)
	fmt.Println("OLS Alpha SE", alphaSe)
	fmt.Println("OLS Beta SE", betaSe)

	fmt.Println("OLS R^2", rSquared)
	fmt.Println("OLS R^2 ADJ", rSquaredAdj)
	fmt.Println("OLS DF", dF)

}
