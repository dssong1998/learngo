package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	urls:=[]string{
		"https://www.google.com",
		"https://www.youtube.com",
		"https://www.amazon.com",
		"https://www.naver.com",
		"https://www.instagram.com",
		"https://www.boogibooks.com",
	}
	c := make(chan string)
	for _, url := range urls {
		go hitUrl(url, c);
	}
	for i := range urls{
		 fmt.Println(<-c, i);
	}
}

func hitUrl(url string, c chan string) error{
	time.Sleep(time.Second)
	fmt.Println("Checking url", url)
	resp, err := http.Get(url);
	if err != nil || resp.StatusCode >= 400 {
		c <- url + " is Failed";
	}
	c <- url + " is OK"
	return nil;
}