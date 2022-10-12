package worker_pool

import "context"

type Worker struct {
	In, Out chan string
}

func (t Worker) Print(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case payload := <-t.In:
			println(payload)
			t.Out <- payload
		}
	}
}

func (t Worker) Check(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case payload := <-t.In:
			println(payload, " checked")
			t.Out <- payload + ":checked"
		}
	}
}
