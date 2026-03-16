package core

import "context"

type Job func() (interface{}, error)

type JobResult struct {
	Result interface{}
	Err    error
}

type jobWrapper struct {
	job    Job
	result chan JobResult
}

type JobQueue struct {
	jobs chan jobWrapper
}

// NewJobQueue creates a single-worker queue for processing vendor DLL tasks synchronously.
func NewJobQueue() *JobQueue {
	return &JobQueue{
		jobs: make(chan jobWrapper, 100),
	}
}

// Start spawns a single goroutine to consume tasks.
func (q *JobQueue) Start(ctx context.Context) {
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case j := <-q.jobs:
				res, err := j.job()
				j.result <- JobResult{Result: res, Err: err}
			}
		}
	}()
}

// Enqueue blocks until the job is executed by the worker and returns the result.
func (q *JobQueue) Enqueue(job Job) (interface{}, error) {
	resCh := make(chan JobResult, 1)
	q.jobs <- jobWrapper{
		job:    job,
		result: resCh,
	}
	res := <-resCh
	return res.Result, res.Err
}
