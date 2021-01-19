package main

import "github.com/go-scrapper/naver_crawler/scrap"

// var columns = map[string]int{
// 	"time":   0,
// 	"price":  1,
// 	"volume": 6,
// }

func main() {
	var stock_codes = []string{"005930", "000250"}
	scrap.UpdateStocks("202101191530", stock_codes)
}
