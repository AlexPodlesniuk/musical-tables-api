package features

type CreateRoom struct {
	Name string
}

func NewCreateRoom(name string) *CreateRoom {
	return &CreateRoom{Name: name}
}
