package main

import (
	"fmt"
	"math"
	"sort"
)

type Excercise1 struct{}

type returnSeries struct {
	year    []float64
	msciRet []float64
	gbiRet  []float64
}

func (e Excercise1) calcMedian(series []float64) float64 {

	// Sort slice ascending
	sort.Slice(series, func(i, j int) bool {
		return series[i] < series[j]
	})

	// Calculate some props
	seriesLen := len(series)
	seriesIsEven := seriesLen%2 == 0
	if seriesIsEven {
		// ret avg on two center pieces
		startP := math.Floor(float64(seriesLen)/float64(2)) - 1
		endP := startP + 1
		fmt.Println(startP, endP)
		return (series[int(startP)] + series[int(endP)]) / 2
	}
	// ret value in center piece
	return series[int(math.Ceil(float64(seriesLen)/float64(2)))-1]

}

func (e Excercise1) proc() {

	series := returnSeries{
		year:    []float64{2000, 2001, 2002, 2003, 2004, 2005, 2006, 2007, 2008, 2009},
		msciRet: []float64{0.4621, -0.0618, 0.0804, 0.2287, 0.4590, 0.2032, 0.4120, -0.0953, -0.1775, -0.4306},
		gbiRet:  []float64{0.1574, -0.034, 0.183, 0.0835, 0.0665, 0.1245, -0.0219, 0.0744, 0.0555, 0.1027},
	}

	// Task 1a)
	// Calculate the arithmetic mean, the geometric mean, and the median of the distribution
	// of the MSCI Germany index.

	N := float64(len(series.year))
	msciSum := float64(0)

	for i := 0; i < int(N); i++ {
		msciElem := series.msciRet[i]
		msciSum += msciElem
	}

	msciArithmeticMean := msciSum / N
	msciGeometricMean := math.Pow(msciSum, (1 / N))
	msciMedian := e.calcMedian(series.msciRet)

	fmt.Println("1a arithmeticMean ", msciArithmeticMean)
	fmt.Println("1a geometricMean ", msciGeometricMean)
	fmt.Println("1a median ", msciMedian)

	// 1b) Describe the dispersion of the distribution of the MSCI Germany index by determining the range, the sample variance, and the sample standard deviation.
	// 1b-1 -- range
	msciCop := make([]float64, len(series.msciRet))
	copy(msciCop, series.msciRet)

	// Sort msciCop asc
	sort.Slice(msciCop, func(i, j int) bool {
		return msciCop[i] < msciCop[j]
	})

	msciMinVal := msciCop[0]
	msciMaxVal := msciCop[len(msciCop)-1]
	msciRange := msciMaxVal - msciMinVal

	// 1b-2 -- the sample variance
	// re-use msciArithmeticMean from before
	msciVariancePool := float64(0)
	for _, val := range msciCop {
		msciVariancePool += math.Pow(val-msciArithmeticMean, 2)
	}
	msciVariance := msciVariancePool / N

	// 1b-3 -- sample standard deviation
	msciStandardDeviation := math.Sqrt(msciVariance)

	fmt.Println("1b range ", msciRange)
	fmt.Println("1b variance ", msciVariance)
	fmt.Println("1b standard deviation ", msciStandardDeviation)

	// Task 1c)
	// To describe the degree to which the distribution of the MSCI Germany index may depart from normality, calculate its skewness and excess kurtosis
	// 1c-1 -- skewness
	msciThirdMomentPool := float64(0)
	msciFourthMomentPool := float64(0)
	for _, val := range msciCop {
		msciThirdMomentPool += math.Pow(val-msciArithmeticMean, 3)  // cubic
		msciFourthMomentPool += math.Pow(val-msciArithmeticMean, 4) // quartic
	}
	msciThirdMoment := msciThirdMomentPool / N
	msciFourthMoment := msciFourthMomentPool / N
	skewnewss := msciThirdMoment / math.Pow(msciStandardDeviation, 3) // third moment / (sd cubic)

	// 1c-2 -- excess kurtosis
	// TODO: Check book if correct formula for kurtosis is used
	kurtosis := msciFourthMoment / math.Pow(msciStandardDeviation, 4) // fourth moment / (sd quartic)

	fmt.Println("1c skewnewss ", skewnewss)
	fmt.Println("1c kurtosis ", kurtosis)

	// todo tmrw:
	// 1d -- Compute the correlation between the MSCI Germany index and the JP Morgan Germany government bond index.
	// 1d-1 caclulate mean, stdev on both series (msci & gbi)
	gbiSum := float64(0)
	for i := 0; i < int(N); i++ {
		gbiElem := series.gbiRet[i]
		gbiSum += gbiElem
	}
	gbiMean := gbiSum / N

	gbiVariancePool := float64(0)
	for _, val := range series.gbiRet {
		gbiVariancePool += math.Pow(val-gbiMean, 2)
	}
	gbiVariance := gbiVariancePool / N
	gbiStandardDeviation := math.Sqrt(gbiVariance)

	// 1d-2 caclulate covariance between msci & gbi: SUM((x - meanX)*(y - meanY))/N
	coVarPool := float64(0)
	for i := 0; i < int(N); i++ {
		gbiElem := series.gbiRet[i]
		msciElem := series.msciRet[i]
		coVarPool += (msciElem - msciArithmeticMean) * (gbiElem - gbiMean)
	}
	coVar := coVarPool / N

	// 1d-3 calculate corr: (COV(x,y)/(SDx * SDy))
	correl := (coVar / (msciStandardDeviation * gbiStandardDeviation))
	fmt.Println("1d msci gbi correlation ", correl)

	// task 1e)
	// Calculate the arithmetic mean return and the standard deviation on a portfolio with
	// 60% invested in the MSCI Germany index and 40% invested in the JP Morgan Germany
	// government bond index.
	// Construct port. slice
	portfolioSlice := []float64{} // make([]float64, len(series.msciRet))

	for i := 0; i < int(N); i++ {
		gbiElem := series.gbiRet[i]
		msciElem := series.msciRet[i]
		combinedRet := (msciElem * 0.6) + (gbiElem * 0.4)
		portfolioSlice = append(portfolioSlice, combinedRet)
	}

	// Calculate mean
	portMeanPool := float64(0)
	for _, v := range portfolioSlice {
		portMeanPool += v
	}

	portMean := portMeanPool / float64(N)

	// Calculate var and stdev
	portVariancePool := float64(0)
	for _, v := range portfolioSlice {
		portVariancePool += math.Pow(v-portMean, 2)
	}
	portVariance := portVariancePool / N
	portStandardDeviation := math.Sqrt(portVariance)

	fmt.Println("1e port mean ", portMean)
	fmt.Println("1e port stdev ", portStandardDeviation)

	//

	//
	// a := object{ x: []int32{0,1} }

}
