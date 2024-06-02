package example

import (
	cqrs_es "cqrs-es"
	"github.com/google/uuid"
)

type Account struct {
	Id        Id
	Owner     string
	Balance   int
	CreatedAt CreationTime
	DeletedAt DeletionTime
	Ver       cqrs_es.Version
}

type Id uuid.UUID

func NewId() Id {
	id, err := uuid.NewV6()
	if err != nil {
		panic(err)
	}
	return Id(id)
}

type CreationTime string

type DeletionTime string

func (a *Account) GetId() Id {
	return a.Id
}

// TODO: Make this impl as default impl for EventSourcedBy interface when Go will allow this.
func (a *Account) ApplyEvent(ev cqrs_es.AppliableEvent[*Account, Id]) {
	ev.ApplyTo(a)
}
