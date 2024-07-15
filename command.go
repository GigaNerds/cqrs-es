package cqrs_es

// TODO: Add support for array of events.

// Command is an interface for domain commands. It is used to produce events to modify pkg.Aggregate.
type Command[Agg Aggregate[ID], ID any, Ev AppliableEvent[Agg, ID]] interface {
	// ExecuteCommand executes a command on a given aggregate object.
	ExecuteCommand(agg Agg) (Ev, error)

	// GetRelatedId returns cqrs_es.Aggregate's ID which is related to this Command.
	GetRelatedId() (ID, error)
}
