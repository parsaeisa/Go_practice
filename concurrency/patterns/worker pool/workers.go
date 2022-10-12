package worker_pool

type Worker struct {
	In, Out chan string
}
