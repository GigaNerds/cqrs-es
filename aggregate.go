package cqrs_es

import "github.com/google/uuid"

// Aggregate is an interface for domain `aggregate`s.
type Aggregate[ID any] interface {
	// GetId returns the Aggregate ID.
	GetId() ID
}

// EventSourcedBy is an interface that describes how event must be applied to the object.
type EventSourcedBy[Agg Aggregate[ID], ID any, Ev AppliableEvent[Agg, ID]] interface {
	// ApplyEvent applies an event to the object.
	ApplyEvent(ev Ev)
}

type CommandHandler[Cmd any] interface {
	HandleCommand(cmd Cmd)
}

// Version of the Aggregate.
// Version is monotonously increasing number that represents the version of the aggregate.
type Version uuid.UUID

// NewVersion generates a new Version.
func NewVersion() Version {
	id, err := uuid.NewV6()
	if err != nil {
		panic(err)
	}
	return Version(id)
}

// VersionedAggregate is a wrapper for an Aggregate to store it's version.
type VersionedAggregate[T Aggregate[ID], ID any] struct {
	// State is the Aggregate object itself.
	State T

	// Ver is the Aggregate's Version. It must increase monotonically to track Aggregate's state
	Ver Version
}
