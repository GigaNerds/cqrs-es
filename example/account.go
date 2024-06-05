package example

import (
	core "cqrs-es"
	"github.com/google/uuid"
)

type Account struct {
	Id        Id
	Owner     Owner
	Balance   Balance
	CreatedAt CreationTime
	DeletedAt DeletionTime
	Ver       core.Version
}

// Id is Account's identifier. It uniquely identifies it.
type Id uuid.UUID

func NewId() Id {
	id, err := uuid.NewV6()
	if err != nil {
		panic(err)
	}
	return Id(id)
}

// Balance is Account's balance value.
type Balance int

// Owner is Account's owner name.
type Owner string

// CreationTime is the time Account was created.
type CreationTime string

// DeletionTime is the time Account was deleted.
type DeletionTime string

func (a *Account) GetId() Id {
	return a.Id
}
