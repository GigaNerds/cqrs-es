package command

import (
	"cqrs-es/pkg"
)

// TODO: Add support for array of events.

// Command is an interface for domain commands. It is used to produce events to modify pkg.Aggregate.
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
