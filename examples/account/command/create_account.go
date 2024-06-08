package command

import (
	"cqrs-es/examples/account/domain"
	"cqrs-es/examples/account/event"
	"time"
)

type CreateAccount struct {
	Owner domain.Owner
}

// TODO: Better error handling. Strings in not the best idea. Maybe newtypes for command errors.

// ExecuteCommand describes logic of applying this command to the examples.Account object.
func (c CreateAccount) ExecuteCommand(_ *domain.Account) (event.AccountCreated, error) {
	created := event.AccountCreated{
		AccountId: domain.NewId(),
		Owner:     c.Owner,
		At:        domain.CreationTime(time.Now().UTC().String()),
	}
	return created, nil
}

func (c CreateAccount) GetRelatedId() domain.Id {
	panic("Don't have any `AccountId`")
}
