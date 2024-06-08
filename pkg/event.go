package pkg

// AppliableEvent is an interface that describes how event must be applied to the object.
type AppliableEvent[Agg Aggregate[ID], ID any] interface {
	// ApplyTo applies an event to the object.
	ApplyTo(agg Agg)

	// GetId returns the Aggregate ID which this event is related to.
	GetId() ID
}

// VersionedEvent is an object that represents an event with its version.
type VersionedEvent[E any] struct {
	// State is the event itself.
	State E
	// Ver is the Version of the event.
	Ver Version
}
