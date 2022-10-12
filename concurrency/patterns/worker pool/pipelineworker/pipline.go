package pipelineworker

import "context"

func NewPipeline(ctx context.Context, numberOfPrintWorkers, numberOfCheckWorkers int) (in, out chan string) {
	inCheck := make(chan string)
	inPrint := make(chan string)
	outFinal := make(chan string)

	for i := 0; i < numberOfPrintWorkers; i++ {
		w := Worker{
			In:  inPrint,
			Out: outFinal,
		}
		go w.Print(ctx)
	}

	for i := 0; i < numberOfCheckWorkers; i++ {
		w := Worker{
			In:  inCheck,
			Out: inPrint,
		}
		go w.Check(ctx)
	}

	return inCheck, outFinal
}
