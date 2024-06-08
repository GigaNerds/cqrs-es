package command

import (
	"cqrs-es/examples/account/domain"
	"cqrs-es/examples/account/event"
	"time"
)

type Deposit struct {
	AccountId domain.AccountId
	Amount    domain.AccountBalance
}

// ExecuteCommand describes logic of applying this command to the examples.Account object.
func (c Deposit) ExecuteCommand(_ *domain.Account) (event.AccountDeposit, error) {
	deposit := event.AccountDeposit{
		AccountId: c.AccountId,
		Amount:    c.Amount,
		At:        event.DepositTime(time.Now().UTC().String()),
	}
	return deposit, nil
}

func (c Deposit) GetRelatedId() domain.AccountId {
	return c.AccountId
}
