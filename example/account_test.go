package example

import (
	"testing"
)

func CreatedAccount() Account {
	acc := Account{}
	event := AccountCreated{
		Owner: "test",
		At:    "2021-01-01 00:00:00",
	}
	acc.ApplyEvent(&event)

	return acc
}

func (a *Account) WithBalance(balance int) {
	a.Balance = balance
}

func TestApplyCreated(t *testing.T) {
	acc := Account{}
	event := AccountCreated{
		Owner: "test",
		At:    "2021-01-01 00:00:00",
	}

	acc.ApplyEvent(&event)

	if acc.Owner != "test" {
		t.Errorf("Expected owner to be 'test', got '%s'", acc.Owner)
	}
	if acc.Balance != 0 {
		t.Errorf("Expected balance to be 0, got '%d'", acc.Balance)
	}
	if acc.DeletedAt != "" {
		t.Errorf("Expected deleted_at to be '', got '%s'", acc.DeletedAt)
	}
	if acc.CreatedAt != "2021-01-01 00:00:00" {
		t.Errorf("Expected created_at to be '2021-01-01 00:00:00', got '%s'", acc.CreatedAt)
	}
}

func TestApplyDeposit(t *testing.T) {
	acc := CreatedAccount()
	event := AccountDeposit{
		AccountId: acc.Id,
		Amount:    10,
		At:        "2022-01-01 00:00:00",
	}

	acc.ApplyEvent(&event)

	if acc.Balance != 10 {
		t.Errorf("Expected balance to be 10, got '%d'", acc.Balance)
	}
}

func TestApplyWithdrawal(t *testing.T) {
	acc := CreatedAccount()
	acc.WithBalance(20)

	event := AccountWithdrawal{
		AccountId: acc.Id,
		Amount:    10,
		At:        "2022-01-01 00:00:00",
	}

	acc.ApplyEvent(&event)

	if acc.Balance != 10 {
		t.Errorf("Expected balance to be 10, got '%d'", acc.Balance)
	}
}
