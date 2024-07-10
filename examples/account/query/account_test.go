package query

import (
	"cqrs-es/examples/account"
	"cqrs-es/examples/account/domain"
	"testing"
	"time"
)

func TestHandleById(t *testing.T) {
	svc := account.NewService()

	id := domain.NewId()
	agg := domain.Account{
		Id:        id,
		Owner:     "test",
		Balance:   20,
		CreatedAt: domain.CreationTime(time.Now().UTC().String()),
		DeletedAt: "",
	}
	svc.Repo.SaveAggregate(&agg)

	q := ById{
		Id: id,
	}
	q.ExecuteQuery(svc)
}
