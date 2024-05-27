package generator

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
)

func FileBuiler(rows []*FileRow) {
	fmt.Println("Exporting data created into: output.csv")
	file, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	if err := gocsv.MarshalFile(&rows, file); err != nil {
		panic(err)
	}
	fmt.Println("Data creation process finished.")
}
