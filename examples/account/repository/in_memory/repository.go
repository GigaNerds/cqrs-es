package in_memory

import (
	"github.com/GigaNerds/cqrs_es"
	"github.com/GigaNerds/cqrs_es/examples/account/domain"
)

// Repository is a repository for account domain.
type Repository struct {
	Accounts map[domain.AccountId]domain.Account

	AccountEvents map[domain.AccountId][]cqrs_es.AppliableEvent[*domain.Account, domain.AccountId]
}

func NewRepository() Repository {
	repo := Repository{
		Accounts:      make(map[domain.AccountId]domain.Account),
		AccountEvents: make(map[domain.AccountId][]cqrs_es.AppliableEvent[*domain.Account, domain.AccountId]),
	}

	return repo
}
