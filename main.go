package main

import (
	"errors"
	"fmt"
	"net/http"
)

var errRequestFailed=errors.New("request failed")

func main(){
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
		result := "OK"
		err := hitURL(url)

		if err !=nil{
			result="FAILED"
		}

		results[url]=result
	}

	for url, res := range results{
		fmt.Println(url, res)
	}
}

func hitURL(url string) error{
	fmt.Println("Checking",url)
	res, err:= http.Get(url)

	if err!=nil|| res.StatusCode>=400{
		fmt.Println(err, res.StatusCode)
		return errRequestFailed
	}
	return nil
}