package main

import (
	"errors"
	"net/http"
)

var errRequestFailed=errors.New("request failed")

func main(){
	urls := []string{
		"https://www.aribnb.com/",
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
		hitURL(url)
	}
}

func hitURL(url string) error{
	res, err:= http.Get(url)

	if err==nil|| res.StatusCode>=400{
		return errRequestFailed
	}
	return nil
}