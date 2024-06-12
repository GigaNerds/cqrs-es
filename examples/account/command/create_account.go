package command

import (
	"cqrs-es/examples/account"
	"cqrs-es/examples/account/domain"
	"cqrs-es/examples/account/event"
	"errors"
	"time"
)

type CreateAccount struct {
	Owner domain.AccountOwner
}

// TODO: Better error handling. Strings in not the best idea. Maybe newtypes for command errors.

// ExecuteCommand describes logic of applying this command to the examples.Account object.
func (c CreateAccount) ExecuteCommand(_ *domain.Account) (account.Event, error) {
	created := event.AccountCreated{
		AccountId: domain.NewId(),
		Owner:     c.Owner,
		At:        domain.CreationTime(time.Now().UTC().String()),
	}
	return &created, nil
}

func (c CreateAccount) GetRelatedId() (domain.AccountId, error) {
	return domain.AccountId{}, errors.New("don't have any `AccountId`")
}
