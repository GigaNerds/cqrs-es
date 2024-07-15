package cqrs_es

// AppliableEvent is an interface that describes how event must be applied to the object.
type AppliableEvent[Agg Aggregate[ID], ID any] interface {
	// ApplyTo applies an event to the object.
	ApplyTo(agg Agg)

	// GetRelatedId returns the Aggregate ID which this event is related to.
	GetRelatedId() ID
}

// VersionedEvent is an object that represents an event with its version.
type VersionedEvent[E any] struct {
	// State is the event itself.
	State E
	// Ver is the Version of the event.
	Ver Version
}

// EventSet is object which is used to combine multiple AppliableEvent's. It's usable to return this object
// from Command with events stored in it. It applies AppliableEvent iterface.
type EventSet[Agg Aggregate[ID], ID comparable, Ev AppliableEvent[Agg, ID]] struct {
	Events []Ev
}

func NewEventSet[Agg Aggregate[ID], ID comparable, Ev AppliableEvent[Agg, ID]](evs []Ev) EventSet[Agg, ID, Ev] {
	return EventSet[Agg, ID, Ev]{
		Events: evs,
	}
}

func (es *EventSet[Agg, Id, Ev]) ApplyTo(agg Agg) {
	for _, ev := range es.Events {
		ev.ApplyTo(agg)
	}
}

func (es *EventSet[Agg, ID, Ev]) GetRelatedId() ID {
	id := es.Events[0].GetRelatedId()
	for _, ev := range es.Events {
		evId := ev.GetRelatedId()
		if evId != id {
			panic("events have incompatible Id's")
		}
	}

	return id
}
