package urlchecker

import (
	"errors"
	"fmt"
	"net/http"
)

var (
	errRequestFailed = errors.New("Request Failed")
)

func hitURL(url string, c chan [2]string) {
	fmt.Println("Checking: ", url)
	resp, err := http.Get(url)
	if err != nil || resp.StatusCode >= 400 {
		fmt.Println(err, resp.StatusCode)
		c <- [2]string{url, "false"}
	} else {
		c <- [2]string{url, "true"}
	}

}

// HitURL hits urls
func HitURL(urls []string, results map[string]string) {
	c := make(chan [2]string)
	n := len(urls)

	for _, url := range urls {
		go hitURL(url, c)
	}

	for i := 0; i < n; i++ {
		// wait in here
		r := <-c
		if r[1] == "true" {
			results[r[0]] = "OK"
		} else {
			results[r[0]] = "FAILED"
		}
	}
}
