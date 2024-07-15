package repository

import cqrs_es "github.com/GigaNerds/cqrs-es"

type AggregateStorage[Agg cqrs_es.Aggregate[ID], ID any] interface {
	SaveAggregate(agg Agg) error

	LoadAggregate(id ID) (Agg, error)
}

type EventStorage[Ev cqrs_es.AppliableEvent[Agg, ID], Agg cqrs_es.Aggregate[ID], ID any] interface {
	SaveEvent(ev Ev) error
}
