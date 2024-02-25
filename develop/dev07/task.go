package main

import (
	"fmt"
	"time"
)

func sig(after time.Duration) <-chan interface{} {
	c := make(chan interface{})
	go func() {
		defer close(c)
		time.Sleep(after)
	}()
	return c
}

func or(channels ...<-chan interface{}) <-chan interface{} {
	resCh := make(chan interface{})

	for _, channel := range channels {
		go func(ch <-chan interface{}) {
			defer close(resCh)
			// когда канал закроется, цикл
			// завершится и первая завершенная горутина
			// закроет single-канал
			for range ch {
				continue
			}
		}(channel)
	}

	return resCh
}

func main() {
	start := time.Now()
	<-or(
		sig(2*time.Hour),
		sig(5*time.Minute),
		sig(1*time.Second),
		sig(1*time.Hour),
		sig(1*time.Minute),
	)
	fmt.Println(time.Since(start))
}
