package features

type CreateRoom struct {
	Name string `validate:"required"`
}

func NewCreateRoom(name string) *CreateRoom {
	return &CreateRoom{Name: name}
}
