package mock

import (
	"sync"
	"time"

	"github.com/ThatsMrTalbot/example/queues"
)

type mailQueuePackage struct {
	queue *mailQueueMock
	item  *queues.MailJob
}

type mailQueueMock struct {
	mutex sync.Mutex

	working map[*queues.MailJob]struct{}
	queue   []*queues.MailJob
	fail    []*queues.MailJob
}

type mailQueueConsumer struct {
	queue *mailQueueMock
}

func (pkg *mailQueuePackage) Ack() error {
	pkg.queue.mutex.Lock()
	defer pkg.queue.mutex.Unlock()

	delete(pkg.queue.working, pkg.item)
	return nil
}

func (pkg *mailQueuePackage) Fail() error {
	pkg.queue.mutex.Lock()
	defer pkg.queue.mutex.Unlock()

	delete(pkg.queue.working, pkg.item)
	pkg.queue.fail = append(pkg.queue.fail, pkg.item)
	return nil
}

func (pkg *mailQueuePackage) Requeue() error {
	pkg.queue.mutex.Lock()
	defer pkg.queue.mutex.Unlock()

	delete(pkg.queue.working, pkg.item)
	pkg.queue.queue = append(pkg.queue.fail, pkg.item)
	return nil
}

// NewMailJobQueueMock Create MailJobQueueMock
func NewMailJobQueueMock() queues.MailQueue {
	return &mailQueueMock{
		working: make(map[*queues.MailJob]struct{}),
	}
}

func (queue *mailQueueMock) Put(job *queues.MailJob) error {
	queue.mutex.Lock()
	defer queue.mutex.Unlock()
	queue.queue = append(queue.queue, job)

	return nil
}

func (queue *mailQueueMock) Consumer() (queues.MailQueueConsumer, error) {
	return &mailQueueConsumer{
		queue: queue,
	}, nil
}

func (consumer *mailQueueConsumer) Get() (*queues.MailJob, error) {
	var working *queues.MailJob

	consumer.queue.mutex.Lock()
	defer consumer.queue.mutex.Unlock()

	for {
		consumer.queue.mutex.Lock()
		defer consumer.queue.mutex.Unlock()

		if len(consumer.queue.queue) != 0 {
			break
		}

		time.Sleep(100 * time.Millisecond)
	}

	working, consumer.queue.queue = consumer.queue.queue[0], consumer.queue.queue[1:]
	consumer.queue.working[working] = struct{}{}

	pkg := &mailQueuePackage{
		queue: consumer.queue,
		item:  working,
	}

	working.Package = pkg

	return working, nil
}
