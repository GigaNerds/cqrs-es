package pkg

type AggregateStorage[Agg Aggregate[ID], ID any] interface {
	SaveAggregate(agg Agg) error

	LoadAggregate(id ID) (Agg, error)
}

type EventStorage[Ev any] interface {
	SaveEvent(ev Ev) error
}
