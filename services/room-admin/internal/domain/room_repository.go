package domain

import "context"

type RoomRepository interface {
	GetRoomByID(ctx context.Context, id string) (*Room, error)
	SaveRoom(ctx context.Context, room *Room) error
}
