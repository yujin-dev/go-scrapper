package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/goquery"
)

func main() {
	check("005930", "2021011316")
}

func check(stock_code string, time string) {
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
	fmt.Println(doc)

	// doc, err := goquery.NewDocumentFromReader(res.Body)
	// checkErr(err)
	// println(doc)
	searchData := doc.Find(".tah")
	fmt.Println(searchData)
	searchData.Each(func(i int, s *goquery.Selection) {
		if i%7 == 0 {
			fmt.Println(i, s.Text())

		}
	})

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
