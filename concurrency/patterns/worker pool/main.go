package main

import (
	"context"

	"pipelineworker/pipelineworker"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	in, out := pipelineworker.NewPipeline(ctx, 2, 3)

	strings := []string{"parsa", "ali", "mohammad"}

	go func(strings []string) {
		for _, s := range strings {
			in <- s
		}
	}(strings)

	for i := 0; i < len(strings); i++ {
		<-out
	}
}
