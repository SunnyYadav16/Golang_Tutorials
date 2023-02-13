package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func main() {

	var endpoints = []string{
		"https://www.google.com",
		//"https://www.fb.com",
		"https://www.twitter.com",
		"https://www.instagram.com",
		"https://www.youtube.com",
	}

	for _, endpoint := range endpoints {
		go getStatusCode(endpoint)
		wg.Add(1)
	}

	wg.Wait()
}

//func greeter(s string) {
//	for i := 0; i < 5; i++ {
//		println(s)
//	}
//}

func getStatusCode(endpoint string) {
	defer wg.Done()

	resp, err := http.Get(endpoint)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("%d status code for %s\n", resp.StatusCode, endpoint)
}
