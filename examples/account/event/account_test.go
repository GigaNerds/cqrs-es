package event

import (
	"testing"

	"github.com/GigaNerds/cqrs-es/examples/account/domain"
)

func CreatedAccount() domain.Account {
	acc := domain.Account{}
	event := AccountCreated{
		Owner: "test",
		At:    "2021-01-01 00:00:00",
	}
	event.ApplyTo(&acc)

	return acc
}

func TestApplyCreated(t *testing.T) {
	acc := domain.Account{}
	event := AccountCreated{
		Owner: "test",
		At:    "2021-01-01 00:00:00",
	}

	event.ApplyTo(&acc)

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

	event.ApplyTo(&acc)

	if acc.Balance != 10 {
		t.Errorf("Expected balance to be 10, got '%d'", acc.Balance)
	}
}

func TestApplyWithdrawal(t *testing.T) {
	acc := CreatedAccount()
	acc.Balance = 20

	event := AccountWithdrawal{
		AccountId: acc.Id,
		Amount:    10,
		At:        "2022-01-01 00:00:00",
	}

	event.ApplyTo(&acc)

	if acc.Balance != 10 {
		t.Errorf("Expected balance to be 10, got '%d'", acc.Balance)
	}
}
