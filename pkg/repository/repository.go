package repository

type AggregateStorage[Agg Aggregate[ID], ID any] interface {
	SaveAggregate(agg Agg) error

	LoadAggregate(id ID) (Agg, error)
}

type EventStorage[Ev AppliableEvent[Agg, ID], Agg Aggregate[ID], ID any] interface {
	SaveEvent(ev Ev) error
}
