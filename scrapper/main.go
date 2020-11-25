package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	baseURL := "https://kr.indeed.com/jobs?q=go&limit=50"

	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		getPage(baseURL, i)
	}
}

func getPages(url string) (pageCount int) {
	res, err := http.Get(url)
	checkErr(err)
	checkCode(res)

	// response의 body는 io 처리된다. (io.Reader type)
	// TODO: python에서의 HTTP.get처리도 with을 사용해야 할까?
	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pageCount = s.Find("a").Length()
	})

	return
}

func getPage(baseURL string, page int) {
	pageUnit := 50
	pageURL := baseURL + "&start=" + strconv.Itoa(page*pageUnit)
	fmt.Println("Requesting ", pageURL)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status: ", res.StatusCode)
	}

}
