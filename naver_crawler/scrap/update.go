package scrap

import (
	"github.com/PuerkitoBio/goquery"
)

func UpdateStocks(time string, stock_codes []string) []extractedData {
	// time 기준 가장 최근 데이터 추출
	var total []extractedData
	c := make(chan extractedData)
	for _, stock_code := range stock_codes {
		go getFirst(stock_code, time, c)
	}
	for i := 0; i < len(stock_codes); i++ {
		extractedStock := <-c
		total = append(total, extractedStock)
	}
	return total
}

func getFirst(stock_code string, time string, c chan<- extractedData) {
	var url = "https://finance.naver.com/item/sise_time.nhn?code=" + stock_code + "&thistime=" + time + "&page=1"
	doc := getRes(url)

	searchData := doc.Find(".tah")
	recentData := extractedData{stock_code: stock_code, date: time[:8]}

	searchData.Each(func(i int, s *goquery.Selection) {
		// fmt.Println(i, s)
		if i < 7 {

			switch i % 7 {
			case 0:
				recentData.time = s.Text()
			case 1:
				recentData.price = s.Text()
			case 6:
				recentData.volume = s.Text()
			}
		}
	})
	// fmt.Println(recentData)
	c <- recentData
}
