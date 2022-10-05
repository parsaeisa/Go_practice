package main

import (
	"fmt"
	"sync"
)

func main() {

	inputChannel1 := addListElementsFromList([]string{"west", "rhythm", "proper", "secret"})
	inputChannel2 := addListElementsFromList([]string{"unusual", "stand", "skill", "solve"})

	chOut := merge(inputChannel1, inputChannel2)

	exit := make(chan struct{})

	go func() {
		for c := range chOut {
			fmt.Println(c)
		}

		close(exit)
	}()

	<-exit
	fmt.Println("All completed .")
}

func merge(channels ...<-chan string) <-chan string {
	wg := sync.WaitGroup{}
	out := make(chan string)

	send := func(ch <-chan string) {
		for n := range ch {
			out <- n
		}

		wg.Done()
	}

	wg.Add(len(channels))

	for _, channel := range channels {
		go send(channel)
	}

	go func() {
		wg.Wait()

		close(out)
	}()

	return out

}

func addListElementsFromList(list []string) <-chan string {

	ch := make(chan string)
	go func() {
		for _, s := range list {
			ch <- s
		}

		close(ch)
	}()

	return ch
}
