package repository

import "github.com/GigaNerds/cqrs-es/pkg"

type AggregateStorage[Agg pkg.Aggregate[ID], ID any] interface {
	SaveAggregate(agg Agg) error

	LoadAggregate(id ID) (Agg, error)
}

type EventStorage[Ev pkg.AppliableEvent[Agg, ID], Agg pkg.Aggregate[ID], ID any] interface {
	SaveEvent(ev Ev) error
}
