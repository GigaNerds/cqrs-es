package command

import (
	account "cqrs-es/example"
	"time"
)

type Deposit struct {
	AccountId account.Id
	Amount    account.Balance
}

// ExecuteCommand describes logic of applying this command to the example.Account object.
func (c Deposit) ExecuteCommand(_ *account.Account) (account.AccountDeposit, error) {
	deposit := account.AccountDeposit{
		AccountId: c.AccountId,
		Amount:    c.Amount,
		At:        account.DepositTime(time.Now().UTC().String()),
	}
	return deposit, nil
}

func (c Deposit) GetRelatedId() account.Id {
	return c.AccountId
}
