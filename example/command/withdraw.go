package command

import (
	account "cqrs-es/example"
	"errors"
	"time"
)

type Withdraw struct {
	AccountId account.Id
	Amount    account.Balance
}

// ExecuteCommand describes logic of applying this command to the example.Account object.
func (c Withdraw) ExecuteCommand(a *account.Account) (account.AccountWithdrawal, error) {
	if a.Balance < c.Amount {
		return account.AccountWithdrawal{}, errors.New("not enough money")
	}

	withdrawal := account.AccountWithdrawal{
		AccountId: c.AccountId,
		Amount:    c.Amount,
		At:        account.WithdrawalTime(time.Now().UTC().String()),
	}
	return withdrawal, nil
}

func (c Withdraw) GetRelatedId() account.Id {
	return c.AccountId
}
