package domain

type RoomRepository interface {
	GetRoomById(id string) (*Room, error)
	SaveRoom(room *Room) error
}
