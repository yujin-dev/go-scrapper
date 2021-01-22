package main

import (
	"strings"

	"github.com/go-scrapper/naver_crawler/scrap"
)

func main() {
	var stock_codes = []string{"005930", "000250"}
	result := scrap.ScrapeStocks("202101191530", stock_codes)
	name := strings.Join(stock_codes, "_")
	scrap.WriteData(name, result)
	// scrap.UpdateStocks("202101191530", stock_codes)
}
