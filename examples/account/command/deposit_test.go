package command

import (
	"testing"
	"time"

	"github.com/GigaNerds/cqrs-es/examples/account"
	"github.com/GigaNerds/cqrs-es/examples/account/domain"
	"github.com/GigaNerds/cqrs-es/examples/account/event"
)

func TestEmitsDeposit(t *testing.T) {
	acc := domain.Account{
		Id:        domain.NewId(),
		Owner:     "test",
		Balance:   20,
		CreatedAt: domain.CreationTime(time.Now().UTC().String()),
		DeletedAt: "",
	}
	cmd := Deposit{
		AccountId: acc.Id,
		Amount:    10,
	}
	ev, _ := cmd.ExecuteCommand(&acc)
	deposit, ok := ev.(*event.AccountDeposit)
	if !ok {
		t.Errorf("Wrong event type returned")
	}

	if deposit.AccountId != cmd.AccountId {
		t.Errorf("Expected AccountId to be '%s', got '%s'", cmd.AccountId, deposit.AccountId)
	}
	if deposit.Amount != cmd.Amount {
		t.Errorf("Expected amount to be '%d', got '%d'", cmd.Amount, deposit.Amount)
	}
}

func TestHandleDeposit(t *testing.T) {
	svc := account.NewService()
	acc := domain.Account{
		Id:        domain.NewId(),
		Owner:     "test",
		Balance:   20,
		CreatedAt: domain.CreationTime(time.Now().UTC().String()),
		DeletedAt: "",
	}
	_ = svc.Repo.SaveAggregate(&acc)

	cmd := Deposit{
		AccountId: acc.Id,
		Amount:    10,
	}

	agg, ev, err := svc.HandleCommand(cmd)
	if err != nil {
		t.Errorf(err.Error())
	}

	deposit, ok := ev.(*event.AccountDeposit)
	if !ok {
		t.Errorf("Wrong event type returned")
	}

	if deposit.AccountId != cmd.AccountId {
		t.Errorf("Expected AccountId to be '%s', got '%s'", cmd.AccountId, deposit.AccountId)
	}
	if deposit.Amount != cmd.Amount {
		t.Errorf("Expected amount to be '%d', got '%d'", cmd.Amount, deposit.Amount)
	}

	resBalance := acc.Balance + cmd.Amount
	if agg.Balance != resBalance {
		t.Errorf("Expected balance to be '%d', got '%d'", resBalance, agg.Balance)
	}
}
