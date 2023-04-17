package features

type CreateRoomRequestDto struct {
	Name string `json:"name"`
}

type CreateRoomResponseDto struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}
