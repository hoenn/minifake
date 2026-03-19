package test

// Mixed embedded interfaces + methods
// Multiple interfaces per file
// Package private types in fakes
// Embedded interface that is not explicitly generated
// Package type returns

type jobName string
type Job struct {
	name jobName
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
	Get(jobName) *Job
	Refresh()
}

// JobPrinter explicitly excluded in generate directive, but is embedded in JobServer.
type JobPrinter interface {
	Print() error
}

type JobServer interface {
	JobList
	JobQueuer
	JobPrinter
	Health() (bool, error)
}
