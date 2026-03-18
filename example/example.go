// nolint
package example

type Job struct {
	name string
}

//go:generate go run github.com/hoenn/minifake ./example.go JobQueuer,JobList,JobServer
type JobQueuer interface {
	Enqueue(j *Job) error
	Dequeue(j *Job) error
	Queue() []*Job
}

type JobList interface {
	Append(j *Job) error
	Remove(j *Job) error
	Get(string) *Job
	Refresh()
}

type JobServer interface {
	JobList
	JobQueuer
}
