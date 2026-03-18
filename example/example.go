// nolint
package example

type Job struct {
	name string
}

//go:generate go run github.com/hoenn/mimic ./example.go JobQueuer,JobList,JobServer
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

// JobPrinter explicitly not included in the generate directive, but is embedded in JobServer.
type JobPrinter interface {
	Print() error
}

type JobServer interface {
	JobList
	JobQueuer
	JobPrinter
	Health() (bool, error)
}
