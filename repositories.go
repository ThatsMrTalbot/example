package example

import (
	"github.com/ThatsMrTalbot/example/repositories"
	"gopkg.in/mgo.v2"
)

// Repositories container
type Repositories struct {
	Config *Config `inject:""`

	QueueItemRepository repositories.QueueItemRepository
}

// Open database and spin up repositories
func (r *Repositories) Open() error {
	var (
		err      error
		session  *mgo.Session
		database *mgo.Database
	)

	if session, err = mgo.Dial(r.Config.Mongo.URL); err != nil {
		return err
	}

	database = session.DB(r.Config.Mongo.Database)

	r.QueueItemRepository = repositories.NewQueueItemRepository(database)

	return nil
}
