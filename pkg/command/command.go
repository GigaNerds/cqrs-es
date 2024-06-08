package command

import (
	"cqrs-es/pkg"
)

// TODO: Add support for array of events.

// Command is an interface for domain commands. It is used to produce events to modify cqrs_es.Aggregate.
//
// T represents a cqrs_es.Aggregate that will be modified by the Command.
// Ev represents an event type that will be produced by the Command.
type Command[Agg pkg.Aggregate[ID], ID any, Ev pkg.AppliableEvent[Agg, ID]] interface {
	// ExecuteCommand executes a command on a given aggregate object.
	ExecuteCommand(agg Agg) (Ev, error)

	// GetAggId returns cqrs_es.Aggregate's ID which is related to this Command.
	GetAggId() ID
}

// TODO: Maybe make part of Command interface.

type Handler[Svc any, Res any] interface {
	HandleWith(svc Svc) (Res, error)
}
