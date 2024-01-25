// Code generated by ministub; DO NOT EDIT.
package example

// FakeJobQueuer implements JobQueuer.
type FakeJobQueuer struct {
	EnqueueStub func(j *Job) error
	DequeueStub func(j *Job) error
	QueueStub   func() []*Job
}

func (fakeImpl *FakeJobQueuer) Enqueue(j *Job) error {
	return fakeImpl.EnqueueStub(j)
}
func (fakeImpl *FakeJobQueuer) Dequeue(j *Job) error {
	return fakeImpl.DequeueStub(j)
}
func (fakeImpl *FakeJobQueuer) Queue() []*Job {
	return fakeImpl.QueueStub()
}

// FakeJobList implements JobList.
type FakeJobList struct {
	AppendStub func(j *Job) error
	RemoveStub func(j *Job) error
	GetStub    func(arg0 string) *Job
}

func (fakeImpl *FakeJobList) Append(j *Job) error {
	return fakeImpl.AppendStub(j)
}
func (fakeImpl *FakeJobList) Remove(j *Job) error {
	return fakeImpl.RemoveStub(j)
}
func (fakeImpl *FakeJobList) Get(arg0 string) *Job {
	return fakeImpl.GetStub(arg0)
}
