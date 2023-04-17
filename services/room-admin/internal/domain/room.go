package domain

type Room struct {
	id     string
	name   string
	tables []Table
}

func NewRoom(id, name string) *Room {
	return &Room{id: id, name: name, tables: []Table{}}
}

func (room Room) ID() string {
	return room.id
}

func (room Room) Name() string {
	return room.name
}

func (room Room) Tables() []Table {
	return room.tables
}

func (room *Room) AddTable(table Table) {
	room.tables = append(room.tables, table)
}
