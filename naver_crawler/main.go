package main

import (
	"github.com/go-scrapper/naver_crawler/scrap"
)

func main() {
	stock_codes := scrap.ReadStockCode("KOSPI")
	// var stock_codes = []string{"005930", "000250"}
	result := scrap.ScrapeStocks("202102101530", stock_codes)
	scrap.WriteData("KOSPI", result)
	// scrap.UpdateStocks("202101191530", stock_codes)
}
