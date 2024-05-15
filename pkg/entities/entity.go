package entities

type IEntity[ID any] interface {
	Id() ID
	SetId(id ID)
}

type Entity[ID any] struct {
	id ID
}

func NewEntity[ID any](id ID) *Entity[ID] {
	return &Entity[ID]{
		id: id,
	}
}

func (e *Entity[ID]) Id() ID {
	return e.id
}

func (e *Entity[ID]) SetId(id ID) {
	e.id = id
}
