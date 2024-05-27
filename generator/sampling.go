package generator

import (
	"fmt"
	"time"

	rng "github.com/leesper/go_rng"
)

func CreateBetaSample(alpha int, betha int, sampleSize int) []float64 {
	if sampleSize == 0 {
		fmt.Println("Sample size not provided, using default (10_000) value")
		sampleSize = 10000
	}

	// Defining Beta distribution generator
	generator := rng.NewBetaGenerator(time.Now().UnixNano())

	// Defining output list
	resultArr := make([]float64, sampleSize)

	for i := 0; i < sampleSize; i++ {
		resultArr[i] = generator.Beta(float64(alpha), float64(betha))
	}

	return resultArr
}
