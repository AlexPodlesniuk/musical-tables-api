package domain

type Table struct {
	id       string
	name     string
	capacity int
}

func NewTable(id, name string, capacity int) *Table {
	return &Table{id: id, name: name, capacity: capacity}
}
