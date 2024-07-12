package account

import (
	"github.com/GigaNerds/cqrs-es/examples/account/domain"
	"github.com/GigaNerds/cqrs-es/examples/account/repository/in_memory"
	"github.com/GigaNerds/cqrs-es/pkg"
)

type Service struct {
	Repo in_memory.Repository
}

func NewService() Service {
	repo := in_memory.NewRepository()

	svc := Service{
		Repo: repo,
	}

	return svc
}

type Command pkg.Command[*domain.Account, domain.AccountId, Event]

type Event pkg.AppliableEvent[*domain.Account, domain.AccountId]

type EventSet struct {
	pkg.EventSet[*domain.Account, domain.AccountId, Event]
}

func NewSet(evs []Event) EventSet {
	return EventSet{
		pkg.NewSet(evs),
	}
}

func (s Service) HandleCommand(cmd Command) (domain.Account, Event, error) {
	repo := s.Repo

	agg := &domain.Account{}
	id, err := cmd.GetRelatedId()
	if err == nil {
		agg, err = repo.LoadAggregate(id)
		if err != nil {
			return domain.Account{}, nil, err
		}
	}

	ev, err := cmd.ExecuteCommand(agg)
	if err != nil {
		return domain.Account{}, nil, err
	}
	ev.ApplyTo(agg)

	err = repo.SaveAggregate(agg)
	if err != nil {
		return domain.Account{}, nil, err
	}
	err = repo.SaveEvent(ev)
	if err != nil {
		return domain.Account{}, nil, err
	}

	return *agg, ev, nil
}
