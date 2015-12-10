package queues

// Package is queued item
type Package interface {
	Ack() error
	Fail() error
	Requeue() error
}
