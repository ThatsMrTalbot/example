package queues

import (
	"encoding/json"

	"github.com/ThatsMrTalbot/example/repositories"
	"github.com/adjust/redismq"
	"github.com/satori/go.uuid"
)

const (
	// MailQueueName is the name of the queue
	MailQueueName = "mailer"
	// MailQueueConsumerName is the name of the queue consumer
	MailQueueConsumerName = "mailer_consumer"
)

// MailJob is the data passed into the queue
type MailJob struct {
	Package

	Email string `json:"email"`
}

type mailJobPackage struct {
	*redismq.Package

	id   string
	repo repositories.QueueItemRepository
}

func (pkg *mailJobPackage) Ack() error {
	err := pkg.repo.Remove(pkg.id)

	if err != nil {
		return err
	}

	return pkg.Package.Ack()
}

// NewMailJob created new mailJob
func NewMailJob(email string) *MailJob {
	mailJob := &MailJob{
		Email: email,
	}
	return mailJob
}

// MailQueue handles mail queue
type MailQueue interface {
	Consumer() (MailQueueConsumer, error)
	Put(*MailJob) error
}

// MailQueueConsumer Retrieves from queue
type MailQueueConsumer interface {
	Get() (*MailJob, error)
}

type mailQueue struct {
	queue *redismq.BufferedQueue
	repo  repositories.QueueItemRepository
}

type mailQueueConsumer struct {
	queue    *mailQueue
	consumer *redismq.Consumer
}

// CreateMailQueue creates new mail job queue
func CreateMailQueue(host string, port string, password string, db int64, repo repositories.QueueItemRepository) MailQueue {
	queue := redismq.CreateBufferedQueue(host, port, password, db, MailQueueName, 20)

	return &mailQueue{
		queue: queue,
		repo:  repo,
	}
}

// SelectMailQueue selects existing mail job queue
func SelectMailQueue(host string, port string, password string, db int64, repo repositories.QueueItemRepository) (MailQueue, error) {
	queue, err := redismq.SelectBufferedQueue(host, port, password, db, MailQueueName, 20)

	if err != nil {
		return nil, err
	}

	return &mailQueue{
		queue: queue,
		repo:  repo,
	}, nil
}

func (q *mailQueue) Consumer() (MailQueueConsumer, error) {
	consumer, err := q.queue.AddConsumer(MailQueueConsumerName)

	if err != nil {
		return nil, err
	}

	return &mailQueueConsumer{
		consumer: consumer,
		queue:    q,
	}, nil
}

func (q *mailQueue) Put(job *MailJob) error {
	id := uuid.NewV1().String()
	data, err := json.Marshal(job)

	if err != nil {
		return err
	}

	err = q.repo.Store(id, data)

	if err != nil {
		return err
	}

	return q.queue.Put(id)
}

func (c *mailQueueConsumer) Get() (*MailJob, error) {
	pkg, err := c.consumer.Get()

	if err != nil {
		return nil, err
	}

	data, err := c.queue.repo.Fetch(pkg.Payload)

	if err != nil {
		pkg.Fail()
		return nil, err
	}

	job := new(MailJob)

	err = json.Unmarshal(data, job)

	if err != nil {
		return nil, err
	}

	job.Package = &mailJobPackage{
		Package: pkg,
		id:      pkg.Payload,
		repo:    c.queue.repo,
	}

	return job, nil
}
