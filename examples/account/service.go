package account

import (
	"cqrs-es/examples/account/domain"
	"cqrs-es/examples/account/repository/in_memory"
	"cqrs-es/pkg"
	"cqrs-es/pkg/command"
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

// Just wow.... very long and I don't know what to do with this

func (s Service) HandleCommand(
	cmd command.Command[*domain.Account, domain.AccountId, pkg.AppliableEvent[*domain.Account, domain.AccountId]],
) (domain.Account, pkg.AppliableEvent[*domain.Account, domain.AccountId], error) {
	repo := s.Repo

	agg := domain.Account{}
	id, err := cmd.GetRelatedId()
	if err == nil {
		agg, err = repo.LoadAggregate(id)
		if err != nil {
			return domain.Account{}, nil, err
		}
	}

	ev, err := cmd.ExecuteCommand(&agg)
	if err != nil {
		return domain.Account{}, nil, err
	}
	ev.ApplyTo(&agg)

	err = repo.SaveAggregate(agg)
	if err != nil {
		return domain.Account{}, nil, err
	}
	err = repo.SaveEvent(ev)
	if err != nil {
		return domain.Account{}, nil, err
	}

	return agg, ev, nil
}
