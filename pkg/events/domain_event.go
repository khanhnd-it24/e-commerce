package events

type IDomainEvent[AggregateRootID any] interface {
	IEvent
	AggregateRootId() AggregateRootID
}

type DomainEvent[AggregateRootID any] struct {
	*Event
	aggregateRootId AggregateRootID
}

func (e *DomainEvent[AggregateRootID]) AggregateRootId() AggregateRootID {
	return e.aggregateRootId
}
