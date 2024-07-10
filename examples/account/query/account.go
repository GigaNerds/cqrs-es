package query

import (
	"cqrs-es/examples/account"
	"cqrs-es/examples/account/domain"
)

type ById struct {
	Id domain.AccountId
}

func (q ById) ExecuteQuery(svc account.Service) (*domain.Account, error) {
	return svc.Repo.LoadAggregate(q.Id)
}
