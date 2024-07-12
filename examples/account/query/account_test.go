package query

import (
	"testing"
	"time"

	"github.com/GigaNerds/cqrs-es/examples/account"
	"github.com/GigaNerds/cqrs-es/examples/account/domain"
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
	res, err := q.ExecuteQuery(svc)
	if err != nil {
		t.Errorf(err.Error())
	}

	if *res != agg {
		t.Errorf("Wrong object type returned")
	}
}

func TestHandleByOwner(t *testing.T) {
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

	q := ByOwner{
		Owner: agg.Owner,
	}
	res, err := q.ExecuteQuery(svc)
	if err != nil {
		t.Errorf(err.Error())
	}

	if *res != agg {
		t.Errorf("Wrong object type returned")
	}
}
