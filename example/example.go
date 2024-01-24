package example

type Job struct {
	name string
}
type JobQueuer interface {
	Enqueue(j *Job) error
	Dequeue(j *Job) error
	Queue() []*Job
}
