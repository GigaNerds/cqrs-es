package in_memory

import (
	"github.com/GigaNerds/cqrs_es/examples/account/domain"
	"github.com/GigaNerds/cqrs_es/examples/account/repository"
)

// Repository is a repository for account domain.
type Repository struct {
	Accounts map[domain.AccountId]domain.Account

	AccountEvents map[domain.AccountId][]repository.StorableEvent
}

func NewRepository() Repository {
	repo := Repository{
		Accounts:      make(map[domain.AccountId]domain.Account),
		AccountEvents: make(map[domain.AccountId][]repository.StorableEvent),
	}

	return repo
}
