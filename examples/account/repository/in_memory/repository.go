package in_memory

import (
	"cqrs-es/examples/account/domain"
	"cqrs-es/pkg"
)

// Repository is a repository for account domain.
type Repository struct {
	Accounts map[domain.AccountId]domain.Account

	AccountEvents map[domain.AccountId][]pkg.AppliableEvent[*domain.Account, domain.AccountId]
}

func NewRepository() Repository {
	repo := Repository{
		Accounts:      make(map[domain.AccountId]domain.Account),
		AccountEvents: make(map[domain.AccountId][]pkg.AppliableEvent[*domain.Account, domain.AccountId]),
	}

	return repo
}
