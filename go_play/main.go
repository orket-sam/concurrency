package main

import (
	"fmt"
	"time"
)

func HeavyTask(message string, n int) <-chan string {
	c := make(chan string)

	go func() {
		count := 0
		for i := 0; i <= 10; {
			c <- fmt.Sprintf("%s %d", message, count)
			time.Sleep(time.Millisecond * time.Duration(n))
			count += 1

		}
	}()
	return c

}

func Fan(input1, input2 <-chan string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i <= 10; {
			c <- <-input1

		}
	}()
	go func() {
		for i := 0; i <= 10; {
			c <- <-input2

		}
	}()

	return c
}

func main() {
	a := HeavyTask("milky chance", 10)
	b := HeavyTask("hozier", 20)
	c := Fan(a, b)
	for i := 0; i <= 10; {
		println(<-c)
	}

}
