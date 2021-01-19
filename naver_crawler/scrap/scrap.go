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
	price      int
	volume     int
}

func Update(time string, stock_code string) {
	var baseURL = "https://finance.naver.com/item/sise_time.nhn?code=" + stock_code + "&thistime=" + time + "&page=1"
	// var data []extractedData
	c := make(chan []extractedData)
	go getUpdate(baseURL, c)
	extracted := <-c
	fmt.Println(extracted)

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

func getUpdate(url string, mainChan chan<- []extractedData) { // 송신
	// var data []extractedData
	// c := make(chan extractedData)
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchData := doc.Find(".tah")
	searchData.Each(func(i int, s *goquery.Selection) {
		fmt.Println(s)
	})

}

// func getPage(page int, url string, mainChan chan<- []extractedData) {
// 	var data []extractedData
// 	c := make(chan extractedData)
// 	pageUrl = url + "&page=" + page
// 	fmt.Printf(pageurl)
// 	res, err := http.Get(pageUrl)
// 	checkErr(err)
// 	checkCode(res)

// 	searchData := doc.Find(".tah")

// 	searchData.Each(func(i int, s *goquery.Selection) {

// 	})
// }
