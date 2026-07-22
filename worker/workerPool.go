package worker

import (
	"GoTTP/context"
	"GoTTP/handler"
)

type Job struct {
	Handler handler.HandlerFunc
	Context *context.Context
	Done    chan struct{}
}

// WorkerPool manages worker goroutines
type WorkerPool struct {
	JobQueue chan *Job
}

// NewWorkerPool creates a pool with maxWorkers
func NewWorkerPool(maxWorkers int) *WorkerPool {
	pool := &WorkerPool{
		JobQueue: make(chan *Job, 1000),
	}

	for i := 0; i < maxWorkers; i++ {
		go pool.worker()
	}

	return pool
}

// worker goroutine: waits for jobs
func (wp *WorkerPool) worker() {
	for job := range wp.JobQueue {
		job.Handler(job.Context)
		close(job.Done)
	}
}

func (wp *WorkerPool) Submit(h handler.HandlerFunc, ctx *context.Context) {
	done := make(chan struct{})
	wp.JobQueue <- &Job{
		Handler: h,
		Context: ctx,
		Done:    done,
	}
	<-done
}
