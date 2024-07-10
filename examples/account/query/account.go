package query

import (
	"cqrs-es/examples/account"
	"cqrs-es/examples/account/domain"
)

// ById is a query to load domain.Account by it's domain.AccountId.
type ById struct {
	// Id of domain.Account to load.
	Id domain.AccountId
}

func (q ById) ExecuteQuery(svc account.Service) (*domain.Account, error) {
	return svc.Repo.LoadAggregate(q.Id)
}
