package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	concurrency := 50
	in := make(chan int)
	done := make(chan []byte)

	go func() {
		i := 0
		for {
			in <- i
			i++
		}
	}()

	for x := 0; x < concurrency; x++ {
		go processWorker(in, x)
	}
	<-done
}

func processWorker(in chan int, worker int) {
	for x := range in {
		t := time.Duration(rand.Intn(4) * int(time.Second))
		time.Sleep(t)
		fmt.Println("worker:", worker, "in ", t, ": ", int(x))
	}
}
