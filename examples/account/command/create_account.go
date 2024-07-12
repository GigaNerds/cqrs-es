package command

import (
	"cqrs-es/examples/account"
	"cqrs-es/examples/account/domain"
	"cqrs-es/examples/account/event"
	"cqrs-es/pkg"
	"errors"
	"time"
)

type CreateAccount struct {
	Owner domain.AccountOwner
}

// TODO: Better error handling. Strings in not the best idea. Maybe newtypes for command errors.

// ExecuteCommand describes logic of applying this command to the examples.Account object.
// example.Account is created with 2 events. Frist is event.AccountCreated which creates
// unactivated example.Account. Second is event.AccountActivated whic activates it.
func (c CreateAccount) ExecuteCommand(_ *domain.Account) (account.Event, error) {
	created := event.AccountCreated{
		AccountId: domain.NewId(),
		Owner:     c.Owner,
		At:        domain.CreationTime(time.Now().UTC().String()),
	}
	activated := event.AccountActivated{
		AccountId: created.AccountId,
		At:        domain.ActivationTime(created.At),
	}

	evSlice := make([]account.Event, 0)
	evSlice = append(evSlice, &created)
	evSlice = append(evSlice, &activated)

	eventSet := account.EventSet{
		EventSet: pkg.NewSet(evSlice),
	}

	return &eventSet, nil
}

func (c CreateAccount) GetRelatedId() (domain.AccountId, error) {
	return domain.AccountId{}, errors.New("don't have any `AccountId`")
}
