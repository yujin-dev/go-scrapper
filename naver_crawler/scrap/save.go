package scrap

import (
	"encoding/csv"
	"os"
)

func WriteData(name string, data []extractedData) {
	file, err := os.Create(name + ".csv")
	checkErr(err)

	w := csv.NewWriter(file)
	defer w.Flush()

	headers := []string{"stock_code", "date", "time", "price", "volume"}
	wErr := w.Write(headers)
	checkErr(wErr)

	for _, row := range data {
		jwErr := w.Write([]string{
			"A" + row.stock_code,
			row.date,
			row.time,
			row.price,
			row.volume,
		})
		checkErr(jwErr)
	}
}
