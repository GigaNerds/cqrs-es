package command

import (
	"cqrs-es/example"
	"errors"
	"time"
)

type Withdraw struct {
	AccountId example.Id
	Amount    example.Balance
}

// ExecuteCommand describes logic of applying this command to the example.Account object.
func (c Withdraw) ExecuteCommand(a *example.Account) (example.AccountWithdrawal, error) {
	if a.Balance < c.Amount {
		return example.AccountWithdrawal{}, errors.New("not enough money")
	}

	withdrawal := example.AccountWithdrawal{
		AccountId: c.AccountId,
		Amount:    c.Amount,
		At:        example.WithdrawalTime(time.Now().UTC().String()),
	}
	return withdrawal, nil
}

func (c Withdraw) GetRelatedId() example.Id {
	return c.AccountId
}
