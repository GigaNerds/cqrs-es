package account

import (
	"github.com/GigaNerds/cqrs_es"
	"github.com/GigaNerds/cqrs_es/examples/account/domain"
	"github.com/GigaNerds/cqrs_es/examples/account/repository"
	"github.com/GigaNerds/cqrs_es/examples/account/repository/in_memory"
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

type Command cqrs_es.Command[*domain.Account, domain.AccountId, Event]

type Event cqrs_es.AppliableEvent[*domain.Account, domain.AccountId]

type EventSet struct {
	cqrs_es.EventSet[*domain.Account, domain.AccountId, Event]
}

func NewSet(evs []Event) EventSet {
	return EventSet{
		cqrs_es.NewEventSet(evs),
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

	set, ok := ev.(*EventSet)
	if ok {
		storable_evs := make([]repository.StorableEvent, 0)
		for _, ev := range set.Events {
			ev, ok := ev.(repository.StorableEvent)
			if !ok {
				panic("unexpected event type")
			}
			storable_evs = append(storable_evs, ev)
		}

		err = repo.SaveEvents(storable_evs)
		if err != nil {
			return domain.Account{}, nil, err
		}
	} else {
		ev, ok := ev.(repository.StorableEvent)
		if !ok {
			panic("unexpected event type")
		}
		err = repo.SaveEvent(ev)
		if err != nil {
			return domain.Account{}, nil, err
		}
	}

	return *agg, ev, nil
}
