package main

import (
	"fmt"
	"net/http"
	"time"
)

type WebRequest struct {
	url      string
	err      error
	response *http.Response
}

var urls = []string{
	"http://www.google.com",
	"http://www.gmail.com",
	"http://stackoverflow.com",
	"http://www.wikipedia.org/",
	"http://www.espncricinfo.com/",
	"https://india.gov.in/",
	"http://www.youtube.1com",
}

// Use gorouitnes to fetch many URLs
// use 1 channel + user defined data structure + bufferred channel
// see the difference by adding/removing go keyword before the immediately invoked function expression

func main() {
	start := time.Now()
	ch := make(chan *WebRequest, len(urls))
	for _, url := range urls {
		go func(url string) {
			fmt.Println("fetching ", url)
			resp, err := http.Get(url)
			if err != nil {
				ch <- &WebRequest{url, err, nil}
			} else {
				resp.Body.Close()
				ch <- &WebRequest{url, err, resp}
			}
		}(url)
	}
	// for content := range ch {
	// 	fmt.Println(content.url, content.response.Status)
	// }
	for i := 0; i < len(urls); i++ {
		content := <-ch
		if content.err == nil {
			fmt.Println(content.url, content.response.Status)
		} else {
			fmt.Println(content.url, content.err.Error())
		}
	}
	fmt.Println("Total duration = ", time.Since(start))
}
