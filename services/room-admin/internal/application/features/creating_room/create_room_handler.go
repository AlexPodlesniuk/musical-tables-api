package features

import (
	"context"
	"musical-tables-api/services/room-admin/internal/domain"

	"github.com/google/uuid"
)

type CreateRoomHandler struct {
	repository domain.RoomRepository
}

func NewCreateRoomHandler(repository domain.RoomRepository) *CreateRoomHandler {
	return &CreateRoomHandler{repository: repository}
}

func (handler *CreateRoomHandler) Handle(ctx context.Context, command *CreateRoom) (*CreateRoomResponseDto, error) {
	id := uuid.New().String()
	room := domain.NewRoom(id, command.Name)

	err := handler.repository.SaveRoom(ctx, room)
	response := &CreateRoomResponseDto{ID: id, Name: command.Name}
	return response, err
}
