package main

import (
	"github.com/go-scrapper/naver_crawler/scrap"
)

func main() {
	scrap.ScrapStockCodes()
	// var stock_codes = []string{"005930", "000250"}
	// scrap.ScrapeStocks("202102101530", stock_codes)
	// name := strings.Join(stock_codes, "_")
	// scrap.WriteData(name, result)
	// scrap.UpdateStocks("202101191530", stock_codes)
}
