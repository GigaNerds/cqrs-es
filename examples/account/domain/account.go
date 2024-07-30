package domain

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id          AccountId
	Owner       AccountOwner
	Balance     AccountBalance
	ActivatedAt ActivationTime
	CreatedAt   CreationTime
	DeletedAt   DeactivationTime
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

// ActivationTime is the time account was activated.
type ActivationTime time.Time

// CreationTime is the time Account was created.
type CreationTime time.Time

// DeactivationTime is the time Account was deactivated. Account's are never deleted to
// save data.
type DeactivationTime time.Time

func (a *Account) GetId() AccountId {
	return a.Id
}
