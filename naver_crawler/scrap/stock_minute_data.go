package scrap

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

func ReadStockCode(market string) []string {
	file, err := os.Open("C:/Users/yujin/go/src/github.com/go-scrapper/naver_crawler/scrap/all_stock_20210216.csv")
	checkErr(err)
	rdr := csv.NewReader(bufio.NewReader(file))
	rows, _ := rdr.ReadAll()
	c := make(chan string)
	stock_code_list := []string{} // array

	for i := 1; i < len(rows); i++ {
		go filterMarket(market, rows[i], c)
	}
	for i := 1; i < len(rows); i++ {
		code := <-c
		if code != "" {
			stock_code_list = append(stock_code_list, code)
		}
	}
	return stock_code_list
}

func filterMarket(market string, row []string, c chan<- string) {
	market_type := row[6]

	if market_type == market {
		c <- row[1]
	} else {
		c <- ""
	}
}

type stockCode struct {
	stock_code  string
	market_type string
	name        string
}

func getKRXpage(url string) *goquery.Document {
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

func extractData() {
	url := "http://data.krx.co.kr/contents/MDC/MDI/mdiLoader/index.cmd?menuId=MDC0201020201"
	doc := getKRXpage(url)
	fmt.Println(doc)
	searchData := doc.Find("CI-GRID")
	fmt.Println(searchData)
}
