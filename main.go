package main

import (
	"fmt"
	"net/http"
)

// var errRequestFailed=errors.New("request failed")

type result struct {
	url string
	status string
}

func main(){
	channel :=make(chan result)
	// results := map[string]string{}
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
}

func hitURL(url string, channel chan<- result) {
	fmt.Println("Checking",url)
	res, err:= http.Get(url)

	status := "OK"

	if err!=nil || res.StatusCode>=400{
		status = "FAILED"
	}
	channel<-result{url:url,status: status}

	
}


/*
func main(){
	channel := make(chan string)

	people := [2]string{"a","b"}

	for _,person :=range people{
		go count(person, channel)
	}

	for i := 0; i<len(people);i++{
		fmt.Println(<-channel)
	}
	// fmt.Println(<-channel)
}

func count(name string, channel chan string){
	for i:=0; i<4;i++{
		fmt.Println(name, i)
		time.Sleep(time.Second)
	}
	channel<-name
}
*/