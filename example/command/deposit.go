package command

import (
	"cqrs-es/example"
	"time"
)

type Deposit struct {
	AccountId example.Id
	Amount    example.Balance
}

// ExecuteCommand describes logic of applying this command to the example.Account object.
func (c Deposit) ExecuteCommand(_ *example.Account) (example.AccountDeposit, error) {
	deposit := example.AccountDeposit{
		AccountId: c.AccountId,
		Amount:    c.Amount,
		At:        example.DepositTime(time.Now().UTC().String()),
	}
	return deposit, nil
}

func (c Deposit) GetRelatedId() example.Id {
	return c.AccountId
}
