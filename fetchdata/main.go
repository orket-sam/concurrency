package main

import (
	"fmt"
	"net/http"
)

func FetchData(url string) <-chan string {

	c := make(chan string)

	go func() {
		req, reqErr := http.NewRequest("GET", url, nil)

		if reqErr != nil {
			c <- fmt.Sprintf("Request error %s %s", url, reqErr.Error())
			return
		}
		client := http.Client{}
		res, resError := client.Do(req)
		if resError != nil {
			c <- fmt.Sprintf("Response error %s %s", url, resError.Error())
			return
		}
		c <- fmt.Sprintf("%s : %s", url, res.Status)
	}()

	return c
}

func Fan(input1, input2 <-chan string) <-chan string {

	c := make(chan string)

	go func() {
		c <- <-input1
		close(c)
	}()

	go func() {
		c <- <-input2

	}()
	return c
}

func main() {
	input1 := FetchData("https://ebmsuite.com")
	input2 := FetchData("http://127.0.0.1.3001")
	c := Fan(input1, input2)
	for i := range c {
		println(i)

	}

}
