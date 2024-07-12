package in_memory

import (
	"errors"

	"github.com/GigaNerds/cqrs-es/examples/account/domain"
	"github.com/GigaNerds/cqrs-es/pkg"
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

type SelectByOwner struct {
	Owner domain.AccountOwner
}

func (q SelectByOwner) ExecuteOpertation(repo Repository) (*domain.Account, error) {
	for _, acc := range repo.Accounts {
		if acc.Owner == q.Owner {
			return &acc, nil
		}
	}

	return &domain.Account{}, errors.New("not found")
}
