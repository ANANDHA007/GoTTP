package worker

import (
	"GoTTP/http"
	"strconv"
)

type Job struct {
	Req    *http.Request
	RespCh chan *http.Response
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
		go pool.worker(i)
	}

	return pool
}

// worker goroutine: waits for jobs
func (wp *WorkerPool) worker(id int) {
	for job := range wp.JobQueue {
		resp := handleRequest(job.Req)
		job.RespCh <- resp
	}
}

// Example request handler logic
func handleRequest(req *http.Request) *http.Response {
	body := []byte("Hello from Worker Pool for " + req.Path)

	resp := &http.Response{
		Status: "200 OK",
		Headers: map[string]string{
			"Content-Length": strconv.Itoa(len(body)),
			"Connection":     "keep-alive",
		},
		Body: body,
	}
	return resp
}
