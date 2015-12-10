package mock

import (
	"fmt"

	"github.com/ThatsMrTalbot/example/repositories"
)

type queueItemRepositoryMock struct {
	items map[string][]byte
}

// NewQueueItemRepositoryMock Creates mock QueueItemRepository
func NewQueueItemRepositoryMock() repositories.QueueItemRepository {
	return &queueItemRepositoryMock{
		items: make(map[string][]byte),
	}
}

func (repo *queueItemRepositoryMock) Store(id string, data []byte) error {
	repo.items[id] = data
	return nil
}

func (repo *queueItemRepositoryMock) Fetch(id string) ([]byte, error) {
	data, ok := repo.items[id]

	if !ok {
		return nil, fmt.Errorf("Id `%s` not found in queue item mock repository", id)
	}

	return data, nil
}

func (repo *queueItemRepositoryMock) Remove(id string) error {
	delete(repo.items, id)
	return nil
}
