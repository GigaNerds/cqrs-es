package command

import (
	account "cqrs-es/example"
	"time"
)

type CreateAccount struct {
	Owner account.Owner
}

// TODO: Better error handling. Strings in not the best idea. Maybe newtypes for command errors.

// ExecuteCommand describes logic of applying this command to the example.Account object.
func (c CreateAccount) ExecuteCommand(_ *account.Account) (account.AccountCreated, error) {
	created := account.AccountCreated{
		AccountId: account.NewId(),
		Owner:     c.Owner,
		At:        account.CreationTime(time.Now().UTC().String()),
	}
	return created, nil
}

func (c CreateAccount) GetRelatedId() account.Id {
	panic("Don't have any `AccountId`")
}
