package main

import (
	"fmt"
	"net/http"
)

func FetchData(url string) <-chan string {

	c := make(chan string)

	go func() {

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			c <- fmt.Sprintf("Request error: %s \n", err.Error())
		}
		client := http.Client{}
		res, responseErr := client.Do(req)
		if responseErr != nil {
			c <- fmt.Sprintf("Response error: %s \n", responseErr.Error())
		} else {
			c <- fmt.Sprintf("%s %d \n", url, res.StatusCode)

		}
		close(c)

	}()

	return c

}

func main() {
	urls := []string{"https://ebmsuite.com", "http://127.0.0.1/3001", "http://127.0.0.1/3002"}

	inputs := []<-chan string{}
	for _, url := range urls {
		inputs = append(inputs, FetchData(url))
	}
	c := Fan(inputs)

	for s := range c {
		println(s)
	}

}

func Fan(inputs []<-chan string) <-chan string {
	c := make(chan string)
	for _, input := range inputs {
		go func(ch <-chan string) {
			c <- <-ch
		}(input)

	}

	return c
}
