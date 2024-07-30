package repository

import cqrs_es "github.com/GigaNerds/cqrs_es"

// AggregateStorage is an interface for storing aggregates in the repository.
type AggregateStorage[Agg cqrs_es.Aggregate[ID], ID any] interface {
	// SaveAggregate saves an aggregate in the repository storage.
	SaveAggregate(agg Agg) error

	// LoadAggregate loads an aggregate from the repository storage.
	LoadAggregate(id ID) (Agg, error)
}

// EventStorage is an interface for storing events in the repository.
type EventStorage[Ev any] interface {
	// SaveEvent saves a single event in the storage.
	SaveEvent(ev Ev) error

	// SaveEvents saves multiple events in the storage.
	SaveEvents(set []Ev) error
}
