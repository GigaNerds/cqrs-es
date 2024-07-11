package query

import (
	"cqrs-es/examples/account"
	"cqrs-es/examples/account/domain"
	"cqrs-es/examples/account/repository/in_memory"
)

// ById is a query to load domain.Account by it's domain.AccountId.
type ById struct {
	// Id of domain.Account to load.
	Id domain.AccountId
}

func (q ById) ExecuteQuery(svc account.Service) (*domain.Account, error) {
	return svc.Repo.LoadAggregate(q.Id)
}

// ByOwner is a query to load domain.Account by it's domain.AccountOwner.
type ByOwner struct {
	// Owner is owner of domain.Account to load.
	Owner domain.AccountOwner
}

// TODO: Think about in_memory usage. Repository must be an abstraction and
//       not implementation must be used in queries.

func (q ByOwner) ExecuteQuery(svc account.Service) (*domain.Account, error) {
	repo_q := in_memory.SelectByOwner{
		Owner: q.Owner,
	}

	return repo_q.ExecuteOpertation(svc.Repo)
}
