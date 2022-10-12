package pipelineworker

import (
	"context"
	"fmt"
)

func NewPipeline(ctx context.Context, numberOfPrintWorkers, numberOfCheckWorkers int) (in, out chan string) {
	inCheck := make(chan string)
	inPrint := make(chan string)
	outFinal := make(chan string)

	for i := 0; i < numberOfPrintWorkers; i++ {
		w := Worker{
			In:  inPrint,
			Out: outFinal,
		}

		go func(iterator int) {
			fmt.Printf("worker %d started working on printing \n", iterator)
			w.Print(ctx)
		}(i)
	}

	for i := 0; i < numberOfCheckWorkers; i++ {
		w := Worker{
			In:  inCheck,
			Out: inPrint,
		}

		go func(i int) {
			fmt.Printf("worker %d started working on checking \n", i)
			w.Check(ctx)
		}(i)
	}

	return inCheck, outFinal
}
