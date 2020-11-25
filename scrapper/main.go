package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type extractedJob struct {
	id       string
	title    string
	location string
	salary   string
	summary  string
}

func main() {
	var jobs []extractedJob
	baseURL := "https://kr.indeed.com/jobs?q=go&limit=50"

	totalPages := getPages(baseURL)
	for i := 0; i < totalPages; i++ {
		extractedJobs := getPage(baseURL, i)
		jobs = append(jobs, extractedJobs...)
	}

	for _, job := range jobs {
		fmt.Println(job)
	}
	// fmt.Println(jobs)
}

func getPage(baseURL string, page int) (jobs []extractedJob) {
	pageUnit := 50
	pageURL := baseURL + "&start=" + strconv.Itoa(page*pageUnit)
	fmt.Println("Requesting ", pageURL)

	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()
	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	searchCards := doc.Find(".jobsearch-SerpJobCard")
	searchCards.Each(func(i int, card *goquery.Selection) {
		job := extractJob(card)
		jobs = append(jobs, job)
	})

	return
}

func extractJob(card *goquery.Selection) extractedJob {
	id, _ := card.Attr("data-jk")
	title := card.Find(".title>a").Text()
	location := card.Find("sjcl").Text()
	salary := cleanString(card.Find(".salaryText").Text())
	summary := cleanString(card.Find(".summary").Text())
	return extractedJob{
		id:       id,
		title:    title,
		location: location,
		salary:   salary,
		summary:  summary,
	}
}

func cleanString(str string) string {
	return strings.Join(strings.Fields(strings.TrimSpace(str)), " ")
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
