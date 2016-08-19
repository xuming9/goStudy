package main

import "fmt"
import "strconv"
import "time"

func main() {
	c := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			c <- "message" + strconv.Itoa(i)
			if i%2 == 0 {
				fmt.Println(time.Second)
				time.Sleep(2 * time.Second)
			}
		}

	}()

	go func() {
		for j := 0; j < 5; j++ {
			c <- "stop"
		}

	}()
	defer close(c)
	for j := 0; j < 10; j++ {
		fmt.Println(<-c)
	}

}
