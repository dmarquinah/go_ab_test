package generator

import (
	"errors"
	"math/rand"
)

// This is an implementation about picking a value from a list of numbers based on a Probability Density Funcion (pdf)
// Similar to the implementation of numpy.random.choice
func Choice[T any](elements []T, size int, replace bool, pdf []float64, rng *rand.Rand) ([]T, error) {
	if !replace && (size > len(elements)) {
		return nil, errors.New("can't sample more than array size without replacements")
	}

	if len(elements) == 0 || len(pdf) == 0 {
		return nil, errors.New("empty list of elements to choose from or empty list of pdf")
	}

	if len(elements) != len(pdf) {
		return nil, errors.New("must provide same number of elements and pdf")
	}

	samples := make([]T, size)
	pdfCopy := make([]float64, len(pdf))
	copy(pdfCopy, pdf)

	for i := 0; i < size; i++ {
		// Cummulative Density Function
		cdf := CDF(pdf)

		if !replace {
			total := cdf[len(cdf)-1]
			for cdfIndex := range cdf {
				cdf[cdfIndex] /= total
			}
		}

		randFloat := rng.Float64()
		sampledIndex := FindIndexFromRight(randFloat, cdf)
		samples[i] = elements[sampledIndex]

		if !replace {
			pdfCopy[sampledIndex] = 0.0
		}
	}

	return samples, nil
}

func CDF(probs []float64) []float64 {
	cdf := make([]float64, len(probs))
	cum := 0.0

	for i := range cdf {
		cum += probs[i]
		cdf[i] = cum
	}

	return cdf
}

func FindIndexFromRight(value float64, cdf []float64) int {
	for i, cumProb := range cdf {
		if cumProb >= value {
			return i
		}
	}

	return len(cdf) - 1
}
