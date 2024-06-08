package domain

import (
	"github.com/google/uuid"
)

type Account struct {
	Id        AccountId
	Owner     AccountOwner
	Balance   AccountBalance
	CreatedAt CreationTime
	DeletedAt DeletionTime
}

// AccountId is Account's identifier. It uniquely identifies it.
type AccountId uuid.UUID

func NewId() AccountId {
	id, err := uuid.NewV6()
	if err != nil {
		panic(err)
	}
	return AccountId(id)
}

// AccountBalance is Account's balance value.
type AccountBalance int

// AccountOwner is Account's owner name.
type AccountOwner string

// CreationTime is the time Account was created.
type CreationTime string

// DeletionTime is the time Account was deleted.
type DeletionTime string

func (a *Account) GetId() AccountId {
	return a.Id
}
