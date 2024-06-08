package command

import (
	"cqrs-es/examples/account/domain"
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

	if ev.AccountId != cmd.AccountId {
		t.Errorf("Expected AccountId to be '%s', got '%s'", cmd.AccountId, ev.AccountId)
	}
	if ev.Amount != cmd.Amount {
		t.Errorf("Expected amount to be '%d', got '%d'", cmd.Amount, ev.Amount)
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
