package memory_persistence

import (
	"musical-tables-api/services/room-admin/internal/domain"
)

type InMemoryRoomRepository struct {
	collection map[string]*domain.Room
}

func NewInMemoryRoomRepository() *InMemoryRoomRepository {
	return &InMemoryRoomRepository{collection: make(map[string]*domain.Room)}
}

func (repo *InMemoryRoomRepository) GetRoomById(id string) (*domain.Room, error) {
	val, ok := repo.collection[id]

	if !ok {
		return nil, domain.ErrRoomNotFound
	}

	return val, nil
}

func (repo *InMemoryRoomRepository) SaveRoom(room *domain.Room) error {
	repo.collection[room.ID()] = room

	return nil
}
