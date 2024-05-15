package entities

type IAggregateRoot[ID any] interface {
	IEntity[ID]
}

type AggregateRoot[ID any] struct {
	*Entity[ID]
}

func NewAggregateRoot[ID any](id ID) *AggregateRoot[ID] {
	return &AggregateRoot[ID]{
		Entity: NewEntity(id),
	}
}
