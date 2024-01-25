package example

type Job struct {
	name string
}
type JobQueuer interface {
	Enqueue(j *Job) error
	Dequeue(j *Job) error
	Queue() []*Job
}

type JobList interface {
	Append(j *Job) error
	Remove(j *Job) error
	Get(string) *Job
}
