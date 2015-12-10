package example

import (
	"github.com/ThatsMrTalbot/example/queues"
)

// Queues container
type Queues struct {
	Config       *Config       `inject:""`
	Repositories *Repositories `inject:""`

	MailJobQueue queues.MailQueue
}

// Open database and spin up repositories
func (q *Queues) Open() error {
	var err error

	q.MailJobQueue, err = queues.SelectMailQueue(q.Config.Redis.Host, q.Config.Redis.Port, q.Config.Redis.Password, q.Config.Redis.Database, q.Repositories.QueueItemRepository)

	if err != nil {
		return err
	}

	return nil
}
