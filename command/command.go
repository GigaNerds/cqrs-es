package command

import core "cqrs-es"

// Command is an interface for domain commands. It is used to produce events to modify cqrs_es.Aggregate.
//
// T represents a cqrs_es.Aggregate that will be modified by the Command.
// Ev represents an event type that will be produced by the Command.
type Command[Agg core.Aggregate[ID], ID any, Ev core.AppliableEvent[Agg, ID]] interface {
	// ExecuteCommand executes a command on a given aggregate object.
	ExecuteCommand(agg Agg) (Ev, error)

	// GetAggId returns cqrs_es.Aggregate's ID which is related to this Command.
	GetAggId() ID
}
