package main

import "fmt"
import "time"

func main() {
	tunnel1 := make(chan string, 10)
	tunnel2 := make(chan string, 10)
	go worker1(tunnel1, tunnel2)
	go worker2(tunnel1, tunnel2)
	time.Sleep(time.Second * 15)
}

func worker1(tunnel1 chan string, tunnel2 chan string) {

	tick1 := time.NewTicker(time.Second * 1).C
	for {
		select {

		case <-tick1:
			tunnel2 <- "ticker1 ticked in worker 1"
		case <-tunnel1:
			fmt.Println(<-tunnel1)

		}
	}

}

func worker2(tunnel1 chan string, tunnel2 chan string) {

	tick2 := time.NewTicker(time.Second * 2).C

	for {
		select {

		case <-tick2:
			tunnel1 <- "ticker2 ticked in worker 2"
		case <-tunnel2:
			fmt.Println(<-tunnel2)

		}
	}

}
