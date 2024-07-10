package in_memory

import (
	"cqrs-es/examples/account/domain"
	"cqrs-es/pkg"
	"errors"
)

func (r *Repository) SaveAggregate(agg *domain.Account) error {
	r.Accounts[agg.Id] = *agg

	return nil
}

func (r *Repository) LoadAggregate(id domain.AccountId) (*domain.Account, error) {
	agg, exist := r.Accounts[id]
	if !exist {
		return &domain.Account{}, errors.New("not exists")
	}

	return &agg, nil
}

func (r *Repository) SaveEvent(ev pkg.AppliableEvent[*domain.Account, domain.AccountId]) error {
	evs := r.AccountEvents[ev.GetRelatedId()]
	evs = append(evs, ev)

	return nil
}
