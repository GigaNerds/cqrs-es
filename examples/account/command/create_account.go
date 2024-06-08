package command

import (
	"cqrs-es/examples/account"
	"cqrs-es/examples/account/domain"
	"cqrs-es/examples/account/event"
	"time"
)

type CreateAccount struct {
	Owner domain.AccountOwner
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

func (c CreateAccount) GetRelatedId() domain.AccountId {
	panic("Don't have any `AccountId`")
}

func (c CreateAccount) HandleWith(svc account.Service) (domain.Account, error) {
	agg := domain.Account{}

	ev, err := c.ExecuteCommand(&agg)
	if err != nil {
		return domain.Account{}, err
	}
	ev.ApplyTo(&agg)

	repo := svc.Repo
	err = repo.SaveAggregate(agg)
	if err != nil {
		return domain.Account{}, err
	}
	err = repo.SaveEvent(&ev)
	if err != nil {
		return domain.Account{}, err
	}

	return agg, nil
}
