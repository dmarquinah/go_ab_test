package generator

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Defining the struct for each row we write in the CSV
// Using snake-case to following existing conventions
type FileRow struct {
	Id                    int     `csv:"id"`
	Timestamp             float64 `csv:"timestamp"`
	Timestamp_day         float64 `csv:"timestamp_day"`
	Hour_of_day           uint8   `csv:"hour_of_day"`
	Minute_of_day         uint16  `csv:"minute_of_day"`
	Day_of_week           uint8   `csv:"day_of_week"`
	Hsm_template_1_result bool    `csv:"template_1_result"`
	Hsm_template_2_result bool    `csv:"template_2_result"`
	Hsm_template_3_result bool    `csv:"template_3_result"`
}

func Generate(sample_size int) {
	// Calls Sampling
	samples := CreateBetaSample(7, 5, sample_size)

	// Calls Data Generator
	rows := GeneratedRows(samples)

	// Ready to be exported
	FileBuiler(rows)
}

func GeneratedRows(samples []float64) []*FileRow {
	fmt.Println("Generating rows of data...")
	rows := make([]*FileRow, len(samples))
	rngSource := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(rngSource)

	for id := 0; id < len(samples); id++ {
		// The generated timestamp is obtained through the Beta Distribution
		// The Beta distribution has a range of [0, 1] and we will use it to represent the fraction of the day (0-24 hours)
		// For example: 6:00 pm --> 18/24 = 0.75
		tsDay := samples[id] * 24

		// Now we generate the hour of the day by taking the integer part
		hourOfDay := uint8(tsDay)

		// And also the minutes
		minuteOfDay := uint16(tsDay * 60)

		// Also generating the day of week randomly
		dayOfWeek := uint8(rand.Intn(5) + 1)

		// Based on previous generated data, creating the timestamp
		ts := tsDay + (float64(dayOfWeek) * 24)

		// First we create conversion rates by random for each HSM
		conv_rates := CreateBetaSample(5, 30, 3)

		// Defining the result of HSM conversion
		results := make([]uint, 3)

		// Select from Converted or Not Converted
		converted := []uint{0, 1}

		for idx, rate := range conv_rates {
			pdf := []float64{1 - rate, rate}
			selected, err := Choice(converted, len(samples), true, pdf, rng)
			results[idx] = selected[0]
			if err != nil {
				log.Fatal(err)
			}
		}

		rows[id] = &FileRow{
			Id:                    id,
			Timestamp:             ts,
			Timestamp_day:         tsDay,
			Hour_of_day:           hourOfDay,
			Minute_of_day:         minuteOfDay,
			Day_of_week:           dayOfWeek,
			Hsm_template_1_result: results[0] == 1,
			Hsm_template_2_result: results[1] == 1,
			Hsm_template_3_result: results[2] == 1,
		}
	}

	fmt.Printf("Generated %d rows in total...\n", len(rows))
	return rows
}
