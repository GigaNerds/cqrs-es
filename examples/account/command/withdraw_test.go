package command

import (
	"cqrs-es/examples/account"
	"cqrs-es/examples/account/domain"
	"cqrs-es/examples/account/event"
	"testing"
	"time"
)

func TestEmitsWithdraw(t *testing.T) {
	acc := domain.Account{
		Id:        domain.NewId(),
		Owner:     "test",
		Balance:   20,
		CreatedAt: domain.CreationTime(time.Now().UTC().String()),
		DeletedAt: "",
	}
	cmd := Withdraw{
		AccountId: acc.Id,
		Amount:    10,
	}
	ev, _ := cmd.ExecuteCommand(&acc)
	withdraw, ok := ev.(*event.AccountWithdrawal)
	if !ok {
		t.Errorf("Wrong event type returned")
	}

	if withdraw.AccountId != cmd.AccountId {
		t.Errorf("Expected AccountId to be '%s', got '%s'", cmd.AccountId, withdraw.AccountId)
	}
	if withdraw.Amount != cmd.Amount {
		t.Errorf("Expected amount to be '%d', got '%d'", cmd.Amount, withdraw.Amount)
	}
}

func TestEmitsErrorWhenNotEnoughBalance(t *testing.T) {
	acc := domain.Account{
		Id:        domain.NewId(),
		Owner:     "test",
		Balance:   7,
		CreatedAt: domain.CreationTime(time.Now().UTC().String()),
		DeletedAt: "",
	}
	cmd := Withdraw{
		AccountId: acc.Id,
		Amount:    10,
	}
	_, err := cmd.ExecuteCommand(&acc)

	if err.Error() != "not enough money" {
		t.Errorf("Expected error")
	}
}

func TestHandleWithdraw(t *testing.T) {
	svc := account.NewService()
	acc := domain.Account{
		Id:        domain.NewId(),
		Owner:     "test",
		Balance:   20,
		CreatedAt: domain.CreationTime(time.Now().UTC().String()),
		DeletedAt: "",
	}
	_ = svc.Repo.SaveAggregate(&acc)

	cmd := Withdraw{
		AccountId: acc.Id,
		Amount:    10,
	}

	agg, ev, err := svc.HandleCommand(cmd)
	if err != nil {
		t.Errorf(err.Error())
	}

	withdraw, ok := ev.(*event.AccountWithdrawal)
	if !ok {
		t.Errorf("Wrong event type returned")
	}

	if withdraw.AccountId != cmd.AccountId {
		t.Errorf("Expected AccountId to be '%s', got '%s'", cmd.AccountId, withdraw.AccountId)
	}
	if withdraw.Amount != cmd.Amount {
		t.Errorf("Expected amount to be '%d', got '%d'", cmd.Amount, withdraw.Amount)
	}

	resBalance := acc.Balance - cmd.Amount
	if agg.Balance != resBalance {
		t.Errorf("Expected balance to be '%d', got '%d'", resBalance, agg.Balance)
	}
}
