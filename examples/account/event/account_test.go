package event

import (
	"testing"
	"time"

	"github.com/GigaNerds/cqrs_es/examples/account/domain"
)

func CreatedAccount() domain.Account {
	acc := domain.Account{}
	at, _ := time.Parse(time.RFC3339, "2021-01-01 00:00:00")

	event := AccountCreated{
		Owner: "test",
		At:    domain.CreationTime(at),
	}
	event.ApplyTo(&acc)

	return acc
}

func TestApplyCreated(t *testing.T) {
	acc := domain.Account{}
	at, _ := time.Parse(time.RFC3339, "2021-01-01 00:00:00")

	event := AccountCreated{
		Owner: "test",
		At:    domain.CreationTime(at),
	}

	event.ApplyTo(&acc)

	if acc.Owner != "test" {
		t.Errorf("Expected owner to be 'test', got '%s'", acc.Owner)
	}
	if acc.Balance != 0 {
		t.Errorf("Expected balance to be 0, got '%d'", acc.Balance)
	}
	if acc.DeletedAt != domain.DeactivationTime(time.Time{}) {
		t.Errorf("Expected deleted_at to be '', got '%v'", acc.DeletedAt)
	}
	if acc.CreatedAt != domain.CreationTime(at) {
		t.Errorf("Expected created_at to be '2021-01-01 00:00:00', got '%v'", acc.CreatedAt)
	}
}

func TestApplyDeposit(t *testing.T) {
	acc := CreatedAccount()
	at, _ := time.Parse(time.RFC3339, "2021-01-01 00:00:00")

	event := AccountDeposit{
		AccountId: acc.Id,
		Amount:    10,
		At:        DepositTime(at),
	}

	event.ApplyTo(&acc)

	if acc.Balance != 10 {
		t.Errorf("Expected balance to be 10, got '%d'", acc.Balance)
	}
}

func TestApplyWithdrawal(t *testing.T) {
	acc := CreatedAccount()
	acc.Balance = 20
	at, _ := time.Parse(time.RFC3339, "2021-01-01 00:00:00")

	event := AccountWithdrawal{
		AccountId: acc.Id,
		Amount:    10,
		At:        WithdrawalTime(at),
	}

	event.ApplyTo(&acc)

	if acc.Balance != 10 {
		t.Errorf("Expected balance to be 10, got '%d'", acc.Balance)
	}
}
