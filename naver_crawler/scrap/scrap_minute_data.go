package scrap

import (
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// 크롤링해서 가져와서 데이터 타입: string
type extractedData struct {
	stock_code string
	date       string
	time       string
	price      string
	volume     string
}

func ScrapeStocks(time string, stock_codes []string) []extractedData {
	var allData []extractedData
	c := make(chan []extractedData)

	for _, stock_code := range stock_codes {
		go getAll(stock_code, time, c)
	}

	for i := 0; i < len(stock_codes); i++ {
		data := <-c
		allData = append(allData, data...)
	}
	return allData
}

func getRes(url string) *goquery.Document {
	req, err := http.NewRequest("GET", url, nil)
	checkErr(err)
	req.Header.Add("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/87.0.4280.141 Safari/537.36")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	res, err := http.DefaultClient.Do(req)
	checkErr(err)
	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromResponse(res)
	checkErr(err)
	return doc
}

func TotalPage(stock_code string, time string) string {
	var baseURL = "https://finance.naver.com/item/sise_time.nhn?code=" + stock_code + "&thistime=" + time + "&page=1"
	doc := getRes(baseURL)
	href, _ := doc.Find(".pgRR").Find("a").Attr("href")
	pages := strings.Split(href, "page=")
	if len(pages) > 1 {
		return strings.TrimSpace(pages[1])
	} else {
		return "0"
	}
}

func getAll(stock_code string, time string, c chan<- []extractedData) { // []extractedData로 안하면 오류
	var stockData []extractedData
	p := TotalPage(stock_code, time)
	pages, _ := strconv.Atoi(p)
	pageChan := make(chan map[string][]string)
	// var totalData []extractedData
	for i := 0; i <= pages; i++ {
		if i > 0 {
			go getPage(strconv.Itoa(i), stock_code, time, pageChan)
		}
	}

	for i := 0; i <= pages; i++ {
		if i > 0 {
			pageData := <-pageChan
			for i := 0; i < len(pageData["stock_code"]); i++ {
				row := extractedData{
					stock_code: pageData["stock_code"][i],
					date:       pageData["date"][i],
					time:       pageData["time"][i],
					price:      pageData["price"][i],
					volume:     pageData["volume"][i],
				}
				stockData = append(stockData, row)
			}
		}
	}
	c <- stockData

}

func getPage(page string, stock_code string, time string, c chan<- map[string][]string) {
	var url = "https://finance.naver.com/item/sise_time.nhn?code=" + stock_code + "&thistime=" + time + "&page=" + page
	doc := getRes(url)
	searchData := doc.Find(".tah")
	data := map[string][]string{
		"stock_code": {},
		"date":       {},
		"time":       {},
		"price":      {},
		"volume":     {},
	}
	searchData.Each(func(i int, s *goquery.Selection) {
		switch i % 7 {
		case 0:
			data["stock_code"] = append(data["stock_code"], stock_code)
			data["date"] = append(data["date"], time[:8])
			data["time"] = append(data["time"], s.Text())
		case 1:
			data["price"] = append(data["price"], s.Text())
		case 6:
			data["volume"] = append(data["volume"], s.Text())
		}
	})
	// fmt.Println(data)
	c <- data
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
