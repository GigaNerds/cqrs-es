package repository

import (
	"time"

	"github.com/GigaNerds/cqrs_es/examples/account/domain"
)

type StorableEvent interface {
	// GetRelatedId returns the Aggregate ID which this event is related to.
	GetRelatedId() domain.AccountId

	// GetEventType returns the type of the event as string.
	GetEventType() string

	// GetHappenedAt returns the time when the event happened.
	GetHappenedAt() time.Time
}
