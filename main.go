package main

import (
	"fmt"
	"net/http"
)


type requestResult struct {
	url string
	status string
}

func main(){
	channel :=make(chan requestResult)
	results := map[string]string{}
	urls := []string{
		"https://www.airbnb.com/",
		"https://www.google.com/",
		"https://www.amazon.com/",
		"https://www.reddit.com/",
		"https://www.google.com/",
		"https://soundcloud.com/",
		"https://www.facebook.com/",
		"https://www.instagram.com/",
		"https://academy.nomadcoders.co/",
	}

	for _,url :=range urls{
		go hitURL(url, channel)
	}

	for i := 0; i<len(urls);i++{
		result := <-channel

		results[result.url]=result.status
	}

	for url, status := range results{
		fmt.Println(url, status)
	}
}

func hitURL(url string, channel chan<- requestResult) {
	res, err:= http.Get(url)

	status := "OK"

	if err!=nil || res.StatusCode>=400{
		status = "FAILED"
	}
	channel<-requestResult{url:url,status: status}

	
}