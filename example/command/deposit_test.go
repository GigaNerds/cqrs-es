package command

import (
	account "cqrs-es/example"
	"testing"
	"time"
)

func TestEmitsDeposit(t *testing.T) {
	acc := account.Account{
		Id:        account.NewId(),
		Owner:     "test",
		Balance:   20,
		CreatedAt: account.CreationTime(time.Now().UTC().String()),
		DeletedAt: "",
	}
	cmd := Deposit{
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
