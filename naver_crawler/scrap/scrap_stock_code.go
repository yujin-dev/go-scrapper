package scrap

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

type stockCode struct {
	stock_code  string
	market_type string
	name        string
}

func ScrapStockCodes() {
	extractData()
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
