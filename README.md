# GO scrapper 

## indeed
indeed에서 구직 관련 정보 크롤링

https://nomadcoders.co/go-for-beginners 참고


## naver_crawler 
naver finance에서 1분 데이터 크롤링 및 csv저장

### 실행

```$ go run main.go```

### main.go
```go
func main() {
	stock_codes := scrap.ReadStockCode("KOSPI")
	result := scrap.ScrapeStocks("202102101530", stock_codes)
	scrap.WriteData("KOSPI", result)
}
```
market type(`KOSPI`) 및 특정 시간(`202102101530`) 해당 market의 전종목 분 데이터를 가져올 수 있음

※ TODO: naver_crawler\scrap\stock_minute_data.go path 수정
