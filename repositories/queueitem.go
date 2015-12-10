package repositories

import (
	"gopkg.in/mgo.v2"
)

const (
	// QueueItemTableName QueueItem table name
	QueueItemTableName = "QueueItem"
)

// QueueItemRepository Repository for queue items
type QueueItemRepository interface {
	Store(id string, data []byte) error
	Fetch(id string) ([]byte, error)
	Remove(id string) error
}

type queueItemRepository struct {
	collection *mgo.Collection
}

// NewQueueItemRepository Create queue item repository
func NewQueueItemRepository(db *mgo.Database) QueueItemRepository {
	return &queueItemRepository{
		db.C(QueueItemTableName),
	}
}

func (repo *queueItemRepository) Store(id string, data []byte) error {
	item := make(map[string]interface{})
	item["_id"] = id
	item["data"] = data
	return repo.collection.Insert(item)
}

func (repo *queueItemRepository) Fetch(id string) ([]byte, error) {
	item := make(map[string]interface{})
	err := repo.collection.FindId(id).One(item)
	return item["data"].([]byte), err
}

func (repo *queueItemRepository) Remove(id string) error {
	return repo.collection.RemoveId(id)
}
