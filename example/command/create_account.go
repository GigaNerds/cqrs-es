package command

import (
	"cqrs-es/example"
	"time"
)

type CreateAccount struct {
	Owner example.Owner
}

// TODO: Better error handling. Strings in not the best idea. Maybe newtypes for command errors.

// ExecuteCommand describes logic of applying this command to the example.Account object.
func (c CreateAccount) ExecuteCommand(_ *example.Account) (example.AccountCreated, error) {
	created := example.AccountCreated{
		AccountId: example.NewId(),
		Owner:     c.Owner,
		At:        example.CreationTime(time.Now().UTC().String()),
	}
	return created, nil
}

func (c CreateAccount) GetRelatedId() example.Id {
	panic("Don't have any `AccountId`")
}
