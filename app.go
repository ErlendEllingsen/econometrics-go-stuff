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

	// -- Regression & OLS --
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

	// -- T-TEST Hypothesis testing --
	tt := TTest{
		ols: ols,
	}

	// Test of Significance
	tStat := tt.calculateTStat(0.5091, 0.2561, 1)
	fmt.Println("tStat", tStat)

	// One sided test of significance (from the book)
	tCritSingle := float64(1.724718243) // t20;5% (cuz one sided)
	rejectStatusOneSided := tt.rejectTestOfSignificanceTestOneSided(tCritSingle, tStat, false)

	// Two sided test of significance (from the book)
	tCritTwoTailed := float64(-2.086) // t20;5% / 2 (~ t;20;2.5%, cuz two sided) from t-table
	rejectStatusTestOfSignificance := tt.rejectTestOfSignificanceTestTwoSided(tCritTwoTailed, tStat)

	// Confidence interval test
	interval := tt.calculateConfidenceInterval(0.5091, 0.2561, tCritTwoTailed)
	rejectStatusConfidenceIntervalTest := tt.rejectConfidenceIntervalTest(interval, 1)
	fmt.Println("Reject (Test of significance) H0 B>=1", rejectStatusOneSided)
	fmt.Println("Reject (Test of significance, two sided) H0 B=1", rejectStatusTestOfSignificance)
	fmt.Println("Confidence interval", interval)
	fmt.Println("Reject (Confidence interval test) H0 B=1", rejectStatusConfidenceIntervalTest)

	// -- F-TEST Hypothesis testing --
	ft := FTest{}
	ftestTstat := ft.calculateTestStat(436.1, 397.2, 2, 144, 4)
	fmt.Println("F-Test, testStat ", ftestTstat)

	// Excercise sheets
	fmt.Println("---- EXCERCISE SHEETS ----")
	ex1 := Excercise1{}
	ex1.proc()
}
