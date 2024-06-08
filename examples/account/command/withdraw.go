package command

import (
	"cqrs-es/examples/account/domain"
	"cqrs-es/examples/account/event"
	"errors"
	"time"
)

type Withdraw struct {
	AccountId domain.AccountId
	Amount    domain.AccountBalance
}

// ExecuteCommand describes logic of applying this command to the examples.Account object.
func (c Withdraw) ExecuteCommand(a *domain.Account) (event.AccountWithdrawal, error) {
	if a.Balance < c.Amount {
		return event.AccountWithdrawal{}, errors.New("not enough money")
	}

	withdrawal := event.AccountWithdrawal{
		AccountId: c.AccountId,
		Amount:    c.Amount,
		At:        event.WithdrawalTime(time.Now().UTC().String()),
	}
	return withdrawal, nil
}

func (c Withdraw) GetRelatedId() (domain.AccountId, error) {
	return c.AccountId, nil
}
