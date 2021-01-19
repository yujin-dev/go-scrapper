package scrap

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type extractedData struct {
	stock_code string
	date       string
	time       string
	price      string
	volume     string
}

var columns = map[int]string{
	0: "time",
	1: "price",
	6: "volume",
}

func UpdateStocks(time string, stock_codes []string) []extractedData {
	// time 기준 가장 최근 데이터 추출
	var total []extractedData
	c := make(chan extractedData)
	for _, stock_code := range stock_codes {
		go getUpdate(stock_code, time, c)
	}
	for i := 0; i < len(stock_codes); i++ {
		extractedStock := <-c
		total = append(total, extractedStock)
	}
	return total
}

// func Scrape(date string, stock_code string) {
// 	var baseURL = "https://finance.naver.com/item/sise_time.nhn?code=" + stock_code + "&thistime=" + date
// 	var data []extractedData
// 	ch := make(chan []extractedData)
// }

func totalPage(url string) int {
	pages := 0
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})
	return pages
}

func getUpdate(stock_code string, time string, c chan<- extractedData) {
	var baseURL = "https://finance.naver.com/item/sise_time.nhn?code=" + stock_code + "&thistime=" + time + "&page=1"
	fmt.Println(baseURL)
	req, err := http.NewRequest("GET", baseURL, nil)
	checkErr(err)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	// defer res.Body.Close()
	res, err := http.DefaultClient.Do(req)
	checkErr(err)
	doc, err := goquery.NewDocumentFromResponse(res)
	checkErr(err)

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
	fmt.Println(recentData)
	c <- recentData
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request Failed: ", res.StatusCode)
	}
}
